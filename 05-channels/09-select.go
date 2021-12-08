package main

import "fmt"

func main() {
	pair := make(chan int)
	unpair := make(chan int)
	exit := make(chan int)

	go shipping(pair, unpair, exit)
	receive2(pair, unpair, exit)
}

func shipping(pair, unpair, exit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			pair <- i
		} else {
			unpair <- i
		}
	}
	close(pair)
	close(unpair)
	exit <- 0
}

func receive2(pair, unpair, exit <-chan int) {
	for {
		select {
		case value := <-pair:
			fmt.Println("pair:", value)
		case value := <-unpair:
			fmt.Println("unpair:", value)
		case value := <-exit:
			fmt.Println("exit:", value)
			return
		}
	}
}
