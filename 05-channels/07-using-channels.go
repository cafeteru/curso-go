package main

import "fmt"

func main() {
	channel := make(chan int)
	go send(channel)
	receive(channel)
}

func send(channel chan<- int) {
	channel <- 1_000
}

func receive(channel <-chan int) {
	fmt.Println(<-channel)
}
