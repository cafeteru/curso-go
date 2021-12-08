package main

import "fmt"

func main() {
	shippingChannel := make(chan<- int, 2) // only send
	shippingChannel <- 42
	shippingChannel <- 43
	receptionChannel := make(<-chan int, 2) // only receive
	fmt.Println("--------------")
	fmt.Printf("%T\n", shippingChannel)
	fmt.Printf("%T\n", receptionChannel)
}
