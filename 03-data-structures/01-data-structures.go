package main

import "fmt"

func main() {
	var intArray [5]int
	fmt.Println(intArray)
	fmt.Println(len(intArray))

	initialArray := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(initialArray)
	fmt.Println(cap(initialArray))
	for index, value := range initialArray {
		fmt.Println(index, value)
	}

	// Create a partial copy
	fmt.Println(initialArray[:2])  // 0 - 1
	fmt.Println(initialArray[3:])  // 3 - 4 - 5
	fmt.Println(initialArray[2:4]) // 2 - 3

	initialArray = append(initialArray, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(initialArray)

	x := make([]int, 10)
	fmt.Println(x)

	ownMap := map[string]int{
		"Batman":   2,
		"Superman": 1,
	}
	fmt.Println(ownMap)
	fmt.Println(ownMap["batman"])

	if v, ok := ownMap["noExists"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("No exists")
	}
}
