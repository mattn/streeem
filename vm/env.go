package vm

import (
	"fmt"
	"reflect"
	"strings"
)

type Env struct {
	name   string
	env    map[string]reflect.Value
	parent *Env
}

func NewEnv() *Env {
	return &Env{env: make(map[string]reflect.Value), parent: nil}
}

func (e *Env) NewEnv() *Env {
	return &Env{env: make(map[string]reflect.Value), parent: e, name: e.name}
}

func (e *Env) Destroy() {
	if e.parent == nil {
		return
	}
	for k, v := range e.parent.env {
		if v.IsValid() && v.Interface() == e {
			delete(e.parent.env, k)
		}
	}
	e.parent = nil
	e.env = nil
}

func (e *Env) SetName(n string) {
	e.name = n
}

func (e *Env) GetName() string {
	return e.name
}

func (e *Env) Get(k string) (reflect.Value, error) {
	for {
		if e.parent == nil {
			v, ok := e.env[k]
			if !ok {
				return NilValue, fmt.Errorf("Undefined symbol '%s'", k)
			}
			return v, nil
		}
		if v, ok := e.env[k]; ok {
			return v, nil
		}
		e = e.parent
	}
	return NilValue, fmt.Errorf("Undefined symbol '%s'", k)
}

func (e *Env) Set(k string, v reflect.Value) error {
	for {
		if e.parent == nil {
			if _, ok := e.env[k]; !ok {
				return fmt.Errorf("Unknown symbol '%s'", k)
			}
			e.env[k] = v
			return nil
		}
		if _, ok := e.env[k]; ok {
			e.env[k] = v
			return nil
		}
		e = e.parent
	}
	return nil
}

func (e *Env) DefineGlobal(k string, v reflect.Value) error {
	global := e
	for global.parent != nil {
		global = global.parent
	}
	return global.Define(k, v)
}

func (e *Env) Define(k string, v reflect.Value) error {
	if strings.Contains(k, ".") {
		return fmt.Errorf("Unknown symbol '%s'", k)
	}
	e.env[k] = v
	return nil
}

func (e *Env) String() string {
	return e.name
}

func (e *Env) Dump() {
	for k, v := range e.env {
		fmt.Printf("%v = %v\n", k, v)
	}
}
