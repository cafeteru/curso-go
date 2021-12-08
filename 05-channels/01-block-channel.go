package main

import "fmt"

func main() {
	channel := make(chan int) // create a channel without buffet (unbuffered channel)
	channel <- 42             // send to channel
	fmt.Println(<-channel)
}
