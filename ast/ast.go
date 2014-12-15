package ast

type Token struct {
	PosImpl // StmtImpl provide Pos() function.
	Tok     int
	Lit     string
}

type Position struct {
	Line   int
	Column int
}

type Pos interface {
	Position() Position
	SetPosition(Position)
}

type PosImpl struct {
	pos Position
}

func (x *PosImpl) Position() Position {
	return x.pos
}

func (x *PosImpl) SetPosition(pos Position) {
	x.pos = pos
}

type Stmt interface {
	Pos
	stmt()
}

type StmtImpl struct {
	PosImpl
}

func (x *StmtImpl) stmt() {}

type ExprStmt struct {
	StmtImpl
	Expr Expr
}

type IfStmt struct {
	StmtImpl
	If     Expr
	Then   []Stmt
	ElseIf []Stmt
	Else   []Stmt
}

type BreakStmt struct {
	StmtImpl
}

type ContinueStmt struct {
	StmtImpl
}

type ReturnStmt struct {
	StmtImpl
	Exprs []Expr
}

type Expr interface {
	Pos
	expr()
}

type ExprImpl struct {
	PosImpl
}

func (x *ExprImpl) expr() {}

type NumberExpr struct {
	ExprImpl
	Lit string
}

type StringExpr struct {
	ExprImpl
	Lit string
}

type ArrayExpr struct {
	ExprImpl
	Exprs []Expr
}

type PairExpr struct {
	ExprImpl
	Key   string
	Value Expr
}

type IdentExpr struct {
	ExprImpl
	Lit string
}

type UnaryExpr struct {
	ExprImpl
	Operator string
	Expr     Expr
}

type ParenExpr struct {
	ExprImpl
	SubExpr Expr
}

type BinOpExpr struct {
	ExprImpl
	Lhs      Expr
	Operator string
	Rhs      Expr
}

type TernaryOpExpr struct {
	ExprImpl
	Expr Expr
	Lhs  Expr
	Rhs  Expr
}

type CallExpr struct {
	ExprImpl
	Func     interface{}
	Name     string
	SubExprs []Expr
}

type LetExpr struct {
	ExprImpl
	Lhs Expr
	Rhs Expr
}

type AssocExpr struct {
	ExprImpl
	Lhs      Expr
	Operator string
	Rhs      Expr
}

type ConstExpr struct {
	ExprImpl
	Value string
}

type FuncExpr struct {
	ExprImpl
	Name   string
	Stmts  []Stmt
	Args   []string
	VarArg bool
}

type ItemExpr struct {
	ExprImpl
	Value Expr
	Index Expr
}

type PipeLine struct {
	StmtImpl
	Exprs []Expr
}
