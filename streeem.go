package main

import (
	"bufio"
	"fmt"
	"github.com/mattn/streeem/parser"
	"github.com/mattn/streeem/vm"
	"io"
	"io/ioutil"
	"os"
	"reflect"
)

type Seq struct {
	vm.IO
	current int
	max     int
}

func (seq *Seq) ReadWrite() (err error) {
	defer func() {
		if e := recover(); e != nil {
			if ee, ok := e.(error); ok {
				err = ee
			}
		}
	}()
	if seq.current < seq.max {
		seq.current += 1
		seq.W <- seq.current
		return nil
	}
	return io.EOF
}

type STDOUT struct {
	vm.IO
}

func (o *STDOUT) ReadWrite() (err error) {
	if o.R == nil {
		return io.EOF
	}
	v, ok := <-o.R
	if ok {
		_, err = fmt.Println(v)
		return err
	}
	return io.EOF
}

type STDIN struct {
	vm.IO
	buf *bufio.Reader
}

func (i *STDIN) ReadWrite() (err error) {
	if i.W == nil {
		return io.EOF
	}
	b, _, err := i.buf.ReadLine()
	i.W <-string(b)
	return err
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage of %s: [file]\n", os.Args[0])
		os.Exit(1)
	}
	b, err := ioutil.ReadFile(os.Args[1])
	scanner := new(parser.Scanner)
	scanner.Init(string(b))
	stmts, err := parser.Parse(scanner)
	if err != nil {
		if e, ok := err.(*parser.Error); ok {
			fmt.Fprintln(os.Stderr, e.Error())
			os.Exit(1)
		}
	}
	env := vm.NewEnv()

	env.Define("seq", reflect.ValueOf(func(max int) *Seq {
		seq := new(Seq)
		seq.max = max
		return seq
	}))

	stdout := new(STDOUT)
	env.Define("STDOUT", reflect.ValueOf(stdout))

	stdin := new(STDIN)
	stdin.buf = bufio.NewReader(os.Stdin)
	env.Define("STDIN", reflect.ValueOf(stdin))

	_, err = vm.Run(stmts, env)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
