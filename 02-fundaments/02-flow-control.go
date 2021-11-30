package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 6; j++ {
			fmt.Println(i, j)
		}
	}

	i := 0
	for {
		i++
		fmt.Println(i)
		if i < 5 {
			continue
		}
		break
	}
	fmt.Println("End")
}
