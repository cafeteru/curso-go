package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			channel <- i
		}
		close(channel) // Close the channel to avoid deadlock
	}()

	for element := range channel {
		fmt.Println(element)
	}
}
