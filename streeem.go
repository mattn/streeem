package main

import (
	"fmt"
	"github.com/mattn/streeem/parser"
	"github.com/mattn/streeem/vm"
	"log"
)

func main() {
	scanner := new(parser.Scanner)
	scanner.Init(`
seq(100) | {|x|
 if x % 15 == 0 {
   "FizzBuzz"
 }
 else if x % 3 == 0 {
   "Fizz"
 }
 else if x % 5 == 0 {
   "Buzz"
 }
 else {
   x
 }
} | STDOUT
	`)
	stmts, err := parser.Parse(scanner)
	if err != nil {
		if e, ok := err.(*parser.Error); ok {
			log.Fatal(e.Error())
		}
	}
	env := vm.NewEnv()

	v, err := vm.Run(stmts, env)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(v.Interface())
}
