package main

import (
	"fmt"
	"github.com/mattn/streeem/parser"
	"log"
)

func main() {
	scanner := new(parser.Scanner)
/*
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
*/
	scanner.Init(`
seq(100)
	`)
	_, err := parser.Parse(scanner)
	if err != nil {
		if e, ok := err.(*parser.Error); ok {
			log.Fatal(e.Error())
		}
	}
	fmt.Println("Syntax OK")
}
