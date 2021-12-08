package main

import "fmt"

func main() {
	channel := make(chan int) // create a channel without buffet (unbuffered channel)
	go func() {
		channel <- 42 // send to channel
	}()
	fmt.Println(<-channel)
}
