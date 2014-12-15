package vm

import (
	"errors"
	"fmt"
	"github.com/mattn/streeem/ast"
	"github.com/mattn/streeem/parser"
	"io"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

var NilValue = reflect.ValueOf((*interface{})(nil))
var TrueValue = reflect.ValueOf(true)
var FalseValue = reflect.ValueOf(false)

var BreakError = errors.New("Unexpected break statement")
var ContinueError = errors.New("Unexpected continue statement")
var ReturnError = errors.New("Unexpected return statement")

type Error struct {
	Message string
	Pos     ast.Position
}

func NewStringError(pos ast.Pos, err string) error {
	return &Error{Message: err, Pos: pos.Position()}
}

func NewErrorf(pos ast.Pos, format string, args ...interface{}) error {
	return &Error{Message: fmt.Sprintf(format, args...), Pos: pos.Position()}
}

func NewError(pos ast.Pos, err error) error {
	if err == nil {
		return nil
	}
	if err == BreakError || err == ContinueError || err == ReturnError {
		return err
	}
	if pe, ok := err.(*parser.Error); ok {
		return pe
	}
	if ee, ok := err.(*Error); ok {
		return ee
	}
	return &Error{Message: err.Error(), Pos: pos.Position()}
}

func (e *Error) Error() string {
	return e.Message
}

type ReadWriter interface {
	Connect(chan interface{}) chan interface{}
	Close()
	ReadWrite() error
}

type IO struct {
	ReadWriter
	R chan interface{}
	W chan interface{}
}

func (io *IO) Connect(r chan interface{}) chan interface{} {
	io.R = r
	io.W = make(chan interface{})
	return io.W
}

func (io *IO) Close() {
	if io.W != nil {
		close(io.W)
	}
}

type FuncIO struct {
	IO
	f reflect.Value
}

func (f *FuncIO) ReadWrite() (err error) {
	defer func() {
		if e := recover(); e != nil {
			if ee, ok := e.(error); ok {
				err = ee
			}
		}
	}()
	v, ok := <-f.R
	if ok {
		rv := f.f.Call([]reflect.Value{reflect.ValueOf(reflect.ValueOf(v))})
		f.W <- rv[0].Interface().(reflect.Value).Interface()
		return nil
	}
	return io.EOF
}

type ArrayIO struct {
	IO
	c int
	a reflect.Value
}

func (a *ArrayIO) ReadWrite() (err error) {
	defer func() {
		if e := recover(); e != nil {
			if ee, ok := e.(error); ok {
				err = ee
			}
		}
	}()
	if a.c < a.a.Len() {
		a.W <- a.a.Index(a.c).Interface()
		a.c++
		return nil
	}
	return io.EOF
}

type Func func(args ...reflect.Value) (reflect.Value, error)

func ToFunc(f Func) reflect.Value {
	return reflect.ValueOf(f)
}

func isNil(v reflect.Value) bool {
	if !v.IsValid() || v.Kind().String() == "unsafe.Pointer" {
		return true
	}
	if (v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr) && v.IsNil() {
		return true
	}
	return false
}

func isNum(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
		return true
	}
	return false
}

func equal(lhsV, rhsV reflect.Value) bool {
	if isNil(lhsV) && isNil(rhsV) {
		return true
	}
	if lhsV.Kind() == reflect.Interface || lhsV.Kind() == reflect.Ptr {
		lhsV = lhsV.Elem()
	}
	if rhsV.Kind() == reflect.Interface || rhsV.Kind() == reflect.Ptr {
		rhsV = rhsV.Elem()
	}
	if !lhsV.IsValid() || !rhsV.IsValid() {
		return true
	}
	if isNum(lhsV) && isNum(rhsV) {
		if rhsV.Type().ConvertibleTo(lhsV.Type()) {
			rhsV = rhsV.Convert(lhsV.Type())
		}
	}
	if lhsV.CanInterface() && rhsV.CanInterface() {
		return reflect.DeepEqual(lhsV.Interface(), rhsV.Interface())
	}
	return reflect.DeepEqual(lhsV, rhsV)
}

func toBool(v reflect.Value) bool {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float() != 0.0
	case reflect.Int, reflect.Int32, reflect.Int64:
		return v.Int() != 0
	case reflect.Bool:
		return v.Bool()
	case reflect.String:
		if v.String() == "true" {
			return true
		}
		if toInt64(v) != 0 {
			return true
		}
	}
	return false
}

func toFloat64(v reflect.Value) float64 {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Int, reflect.Int32, reflect.Int64:
		return float64(v.Int())
	}
	return 0.0
}

func toString(v reflect.Value) string {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.String {
		return v.String()
	}
	if !v.IsValid() {
		return "nil"
	}
	return fmt.Sprint(v.Interface())
}

func toInt64(v reflect.Value) int64 {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return int64(v.Float())
	case reflect.Int, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.String:
		s := v.String()
		var i int64
		var err error
		if strings.HasPrefix(s, "0x") {
			i, err = strconv.ParseInt(s, 16, 64)
		} else {
			i, err = strconv.ParseInt(s, 10, 64)
		}
		if err == nil {
			return int64(i)
		}
	}
	return 0
}

func Run(stmts []ast.Stmt, env *Env) (reflect.Value, error) {
	rv := NilValue
	var err error
	for _, stmt := range stmts {
		if _, ok := stmt.(*ast.ContinueStmt); ok {
			return NilValue, ContinueError
		}
		rv, err = RunSingleStmt(stmt, env)
		if _, ok := stmt.(*ast.ReturnStmt); ok {
			return reflect.ValueOf(rv), ReturnError
		}
		if err != nil {
			return rv, NewError(stmt, err)
		}
	}
	return rv, nil
}

func pipeLine(stmt *ast.PipeLine, env *Env) (err error) {
	defer func() {
		if e := recover(); e != nil {
			if ee, ok := e.(error); ok {
				err = ee
			}
		}
	}()
	pl := []ReadWriter{}
	var oc chan interface{}
	for _, item := range stmt.Exprs {
		var rw ReadWriter
		rv, err := invokeExpr(item, env)
		if err != nil {
			return NewError(stmt, err)
		}
		switch rv.Kind() {
		case reflect.Func:
			rw = &FuncIO{IO{}, rv}
		case reflect.Slice:
			rw = &ArrayIO{IO{}, 0, rv}
		default:
			if iv, ok := rv.Interface().(ReadWriter); ok {
				rw = iv
			} else {
				return NewError(stmt, errors.New("invalid ReadWriter"))
			}
		}
		oc = rw.Connect(oc)
		pl = append(pl, rw)
	}

	var wg sync.WaitGroup
	for _, rw := range pl {
		wg.Add(1)
		var ge error
		go func(item ReadWriter) {
			defer wg.Done()
			for {
				err = item.ReadWrite()
				if err != nil {
					if err != io.EOF {
						ge = err
					}
					item.Close()
					break
				}
			}
		}(rw)
	}
	wg.Wait()
	return nil
}

func RunSingleStmt(stmt ast.Stmt, env *Env) (reflect.Value, error) {
	switch stmt := stmt.(type) {
	case *ast.ExprStmt:
		rv, err := invokeExpr(stmt.Expr, env)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		return rv, nil
	case *ast.IfStmt:
		// If
		rv, err := invokeExpr(stmt.If, env)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		if toBool(rv) {
			// Then
			newenv := env.NewEnv()
			defer newenv.Destroy()
			rv, err = Run(stmt.Then, newenv)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			return rv, nil
		}
		done := false
		if len(stmt.ElseIf) > 0 {
			for _, stmt := range stmt.ElseIf {
				stmt_if := stmt.(*ast.IfStmt)
				// ElseIf
				rv, err = invokeExpr(stmt_if.If, env)
				if err != nil {
					return rv, NewError(stmt, err)
				}
				if !toBool(rv) {
					continue
				}
				// ElseIf Then
				done = true
				rv, err = Run(stmt_if.Then, env)
				if err != nil {
					return rv, NewError(stmt, err)
				}
				break
			}
		}
		if !done && len(stmt.Else) > 0 {
			// Else
			newenv := env.NewEnv()
			defer newenv.Destroy()
			rv, err = Run(stmt.Else, newenv)
			if err != nil {
				return rv, NewError(stmt, err)
			}
		}
		return rv, nil
	case *ast.ReturnStmt:
		rvs := []interface{}{}
		switch len(stmt.Exprs) {
		case 0:
			return NilValue, nil
		case 1:
			rv, err := invokeExpr(stmt.Exprs[0], env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			return rv, nil
		}
		for _, expr := range stmt.Exprs {
			rv, err := invokeExpr(expr, env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			if isNil(rv) {
				rvs = append(rvs, nil)
			} else if rv.IsValid() {
				rvs = append(rvs, rv.Interface())
			} else {
				rvs = append(rvs, nil)
			}
		}
		return reflect.ValueOf(rvs), nil
	case *ast.PipeLine:
		err := pipeLine(stmt, env)
		return reflect.ValueOf(nil), err
	default:
		return NilValue, NewStringError(stmt, "Unknown statement")
	}
}

func invokeExpr(expr ast.Expr, env *Env) (reflect.Value, error) {
	switch e := expr.(type) {
	case *ast.NumberExpr:
		if strings.Contains(e.Lit, ".") {
			v, err := strconv.ParseFloat(e.Lit, 64)
			if err != nil {
				return NilValue, NewError(expr, err)
			}
			return reflect.ValueOf(float64(v)), nil
		}
		var i int64
		var err error
		if strings.HasPrefix(e.Lit, "0x") {
			i, err = strconv.ParseInt(e.Lit[2:], 16, 64)
		} else {
			i, err = strconv.ParseInt(e.Lit, 10, 64)
		}
		if err != nil {
			return NilValue, NewError(expr, err)
		}
		return reflect.ValueOf(i), nil
	case *ast.IdentExpr:
		return env.Get(e.Lit)
	case *ast.StringExpr:
		return reflect.ValueOf(e.Lit), nil
	case *ast.ArrayExpr:
		a := make([]interface{}, len(e.Exprs))
		for i, expr := range e.Exprs {
			arg, err := invokeExpr(expr, env)
			if err != nil {
				return arg, NewError(expr, err)
			}
			a[i] = arg.Interface()
		}
		return reflect.ValueOf(a), nil
	case *ast.UnaryExpr:
		v, err := invokeExpr(e.Expr, env)
		if err != nil {
			return v, NewError(expr, err)
		}
		switch e.Operator {
		case "-":
			if v.Kind() == reflect.Float64 {
				return reflect.ValueOf(-v.Float()), nil
			}
			return reflect.ValueOf(-v.Int()), nil
		case "^":
			return reflect.ValueOf(^toInt64(v)), nil
		case "!":
			return reflect.ValueOf(!toBool(v)), nil
		default:
			return NilValue, NewStringError(e, "Unknown operator ''")
		}
	case *ast.ParenExpr:
		v, err := invokeExpr(e.SubExpr, env)
		if err != nil {
			return v, NewError(expr, err)
		}
		return v, nil
	case *ast.FuncExpr:
		f := reflect.ValueOf(func(expr *ast.FuncExpr, env *Env) Func {
			return func(args ...reflect.Value) (reflect.Value, error) {
				if !expr.VarArg {
					if len(args) != len(expr.Args) {
						return NilValue, NewStringError(expr, "Arguments Number of mismatch")
					}
				}
				newenv := env.NewEnv()
				if expr.VarArg {
					newenv.Define(expr.Args[0], reflect.ValueOf(args))
				} else {
					for i, arg := range expr.Args {
						newenv.Define(arg, args[i])
					}
				}
				rr, err := Run(expr.Stmts, newenv)
				if err == ReturnError {
					err = nil
					rr = rr.Interface().(reflect.Value)
				}
				return rr, err
			}
		}(e, env))
		env.Define(e.Name, f)
		return f, nil
	case *ast.ItemExpr:
		v, err := invokeExpr(e.Value, env)
		if err != nil {
			return v, NewError(expr, err)
		}
		i, err := invokeExpr(e.Index, env)
		if err != nil {
			return i, NewError(expr, err)
		}
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
			if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
				return NilValue, NewStringError(expr, "Array index should be int")
			}
			ii := int(i.Int())
			if ii < 0 || ii >= v.Len() {
				return NilValue, nil
			}
			return v.Index(ii), nil
		}
		if v.Kind() == reflect.Map {
			if i.Kind() != reflect.String {
				return NilValue, NewStringError(expr, "Map key should be string")
			}
			return v.MapIndex(i), nil
		}
		if v.Kind() == reflect.String {
			rs := []rune(v.Interface().(string))
			ii := int(i.Int())
			if ii < 0 || ii >= len(rs) {
				return NilValue, nil
			}
			return reflect.ValueOf(rs[ii]), nil
		}
		return v, NewStringError(expr, "Invalid operation")
	case *ast.BinOpExpr:
		lhsV := NilValue
		rhsV := NilValue
		var err error

		lhsV, err = invokeExpr(e.Lhs, env)
		if err != nil {
			return lhsV, NewError(expr, err)
		}
		if lhsV.Kind() == reflect.Interface {
			lhsV = lhsV.Elem()
		}
		if e.Rhs != nil {
			rhsV, err = invokeExpr(e.Rhs, env)
			if err != nil {
				return rhsV, NewError(expr, err)
			}
			if rhsV.Kind() == reflect.Interface {
				rhsV = rhsV.Elem()
			}
		}
		switch e.Operator {
		case "+":
			if lhsV.Kind() == reflect.String || rhsV.Kind() == reflect.String {
				return reflect.ValueOf(toString(lhsV) + toString(rhsV)), nil
			}
			if (lhsV.Kind() == reflect.Array || lhsV.Kind() == reflect.Slice) && (rhsV.Kind() != reflect.Array && rhsV.Kind() != reflect.Slice) {
				return reflect.Append(lhsV, rhsV), nil
			}
			if (lhsV.Kind() == reflect.Array || lhsV.Kind() == reflect.Slice) && (rhsV.Kind() == reflect.Array || rhsV.Kind() == reflect.Slice) {
				return reflect.AppendSlice(lhsV, rhsV), nil
			}
			if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
				return reflect.ValueOf(toFloat64(lhsV) + toFloat64(rhsV)), nil
			}
			return reflect.ValueOf(toInt64(lhsV) + toInt64(rhsV)), nil
		case "-":
			if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
				return reflect.ValueOf(toFloat64(lhsV) - toFloat64(rhsV)), nil
			}
			return reflect.ValueOf(toInt64(lhsV) - toInt64(rhsV)), nil
		case "*":
			if lhsV.Kind() == reflect.String && (rhsV.Kind() == reflect.Int || rhsV.Kind() == reflect.Int32 || rhsV.Kind() == reflect.Int64) {
				return reflect.ValueOf(strings.Repeat(toString(lhsV), int(toInt64(rhsV)))), nil
			}
			if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
				return reflect.ValueOf(toFloat64(lhsV) * toFloat64(rhsV)), nil
			}
			return reflect.ValueOf(toInt64(lhsV) * toInt64(rhsV)), nil
		case "/":
			return reflect.ValueOf(toFloat64(lhsV) / toFloat64(rhsV)), nil
		case "%":
			return reflect.ValueOf(toInt64(lhsV) % toInt64(rhsV)), nil
		case "==":
			return reflect.ValueOf(equal(lhsV, rhsV)), nil
		case "!=":
			return reflect.ValueOf(equal(lhsV, rhsV) == false), nil
		case ">":
			return reflect.ValueOf(toFloat64(lhsV) > toFloat64(rhsV)), nil
		case ">=":
			return reflect.ValueOf(toFloat64(lhsV) >= toFloat64(rhsV)), nil
		case "<":
			return reflect.ValueOf(toFloat64(lhsV) < toFloat64(rhsV)), nil
		case "<=":
			return reflect.ValueOf(toFloat64(lhsV) <= toFloat64(rhsV)), nil
		case "|":
			return reflect.ValueOf(toInt64(lhsV) | toInt64(rhsV)), nil
		case "||":
			if toBool(lhsV) {
				return lhsV, nil
			}
			return rhsV, nil
		case "&":
			return reflect.ValueOf(toInt64(lhsV) & toInt64(rhsV)), nil
		case "&&":
			if toBool(lhsV) {
				return rhsV, nil
			}
			return lhsV, nil
		case "**":
			if lhsV.Kind() == reflect.Float64 {
				return reflect.ValueOf(math.Pow(toFloat64(lhsV), toFloat64(rhsV))), nil
			}
			return reflect.ValueOf(int64(math.Pow(toFloat64(lhsV), toFloat64(rhsV)))), nil
		case ">>":
			return reflect.ValueOf(toInt64(lhsV) >> uint64(toInt64(rhsV))), nil
		case "<<":
			return reflect.ValueOf(toInt64(lhsV) << uint64(toInt64(rhsV))), nil
		default:
			return NilValue, NewStringError(expr, "Unknown operator")
		}
	case *ast.ConstExpr:
		switch e.Value {
		case "true":
			return reflect.ValueOf(true), nil
		case "false":
			return reflect.ValueOf(false), nil
		}
		return reflect.ValueOf(nil), nil
	case *ast.CallExpr:
		f := NilValue

		if e.Func != nil {
			f = e.Func.(reflect.Value)
		} else {
			var err error
			ff, err := env.Get(e.Name)
			if err != nil {
				return f, err
			}
			f = ff
		}
		_, isReflect := f.Interface().(Func)

		args := []reflect.Value{}
		for i, expr := range e.SubExprs {
			arg, err := invokeExpr(expr, env)
			if err != nil {
				return arg, NewError(expr, err)
			}

			if i < f.Type().NumIn() {
				if !f.Type().IsVariadic() {
					it := f.Type().In(i)
					if arg.Kind().String() == "unsafe.Pointer" {
						arg = reflect.New(it).Elem()
					}
					if arg.Kind() != it.Kind() && arg.IsValid() && arg.Type().ConvertibleTo(it) {
						arg = arg.Convert(it)
					} else if !arg.IsValid() {
						arg = reflect.Zero(it)
					}
				}
			}
			if !arg.IsValid() {
				arg = NilValue
			}

			if !isReflect {
				args = append(args, arg)
			} else {
				if arg.Kind() == reflect.Interface {
					arg = arg.Elem()
				}
				args = append(args, reflect.ValueOf(arg))
			}
		}
		ret := NilValue
		var err error
		func() {
			defer func() {
				if os.Getenv("ANKO_DEBUG") == "" {
					if ex := recover(); ex != nil {
						if e, ok := ex.(error); ok {
							err = e
						} else {
							err = errors.New(fmt.Sprint(ex))
						}
					}
				}
			}()
			if f.Kind() == reflect.Interface {
				f = f.Elem()
			}
			rets := f.Call(args)
			if isReflect {
				ev := rets[1].Interface()
				if ev != nil {
					err = ev.(error)
				}
				ret = rets[0].Interface().(reflect.Value)
			} else {
				if f.Type().NumOut() == 1 {
					ret = rets[0]
				} else {
					var result []interface{}
					for _, r := range rets {
						result = append(result, r.Interface())
					}
					ret = reflect.ValueOf(result)
				}
			}
		}()
		if err != nil {
			return ret, NewError(expr, err)
		}
		return ret, nil
	case *ast.TernaryOpExpr:
		rv, err := invokeExpr(e.Expr, env)
		if err != nil {
			return rv, NewError(expr, err)
		}
		if toBool(rv) {
			lhsV, err := invokeExpr(e.Lhs, env)
			if err != nil {
				return lhsV, NewError(expr, err)
			}
			return lhsV, nil
		}
		rhsV, err := invokeExpr(e.Rhs, env)
		if err != nil {
			return rhsV, NewError(expr, err)
		}
		return rhsV, nil
	default:
		return NilValue, NewStringError(expr, "Unknown expression")
	}
}
