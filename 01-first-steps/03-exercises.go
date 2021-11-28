package main

import "fmt"

type MyNumber int

var x int
var y string
var z bool
var a MyNumber

func main() {
	fmt.Println("######### Exercise 01 #########")
	exercise01()
	fmt.Println("######### Exercise 02 #########")
	exercise02()
	fmt.Println("######### Exercise 03 #########")
	exercise03()
	fmt.Println("######### Exercise 04 #########")
	exercise04()
}

func exercise01() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(x, y, z)
	fmt.Printf("%d, \"%s\", %t\n", x, y, z)
}

func exercise02() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func exercise03() {
	x = 42
	y = "James Bond"
	z = true

	fmt.Printf("%d, \"%s\", %t\n", x, y, z)
}

func exercise04() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	a := 83
	fmt.Println(a)
}
