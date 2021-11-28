package main

import "fmt"

type Money int

var intVariable int
var moneyVariable Money

func main() {
	intVariable = 42
	fmt.Println(intVariable)
	fmt.Printf("%T\n", intVariable)

	moneyVariable = 1000
	fmt.Println(moneyVariable)
	fmt.Printf("%T\n", moneyVariable)

	intVariable = int(moneyVariable) // conversion, no cast
	fmt.Println(intVariable)
	fmt.Printf("%T\n", intVariable)
}
