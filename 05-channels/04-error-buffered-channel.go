package main

import "fmt"

func main() {
	channel := make(chan int, 1) // create a channel with buffet (buffered channel)
	channel <- 42                // send to channel
	channel <- 43                // block code
	fmt.Println(<-channel)
}
