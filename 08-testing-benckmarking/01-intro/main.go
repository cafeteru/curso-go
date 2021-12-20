package main

import "fmt"

func main() {
	var total = Sum(1, 2, 3, 4)
	fmt.Println(total)
}

func Sum(numbers ...int) int {
	var sum = 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}
