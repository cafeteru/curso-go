// Package main first package of program
package main

import "fmt"

func main() {
	var total = Sum(1, 2, 3, 4)
	fmt.Println(total)
}

// Sum sum all numbers
func Sum(numbers ...int) int {
	var sum = 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}
