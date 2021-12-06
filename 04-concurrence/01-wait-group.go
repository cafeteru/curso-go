package main

import (
	"fmt"
	"runtime"
	"sync"
)

var limit = 10
var waitGroup sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("GO routines\t", runtime.NumGoroutine())

	waitGroup.Add(1) // add a thread
	go foo()         // go -> launch in another thread
	fmt.Println("GO routines\t", runtime.NumGoroutine())
	bar()
	waitGroup.Wait() // Wait for all threads to finish
}

func foo() {
	for i := 0; i < limit; i++ {
		fmt.Println("foo: ", i)
	}
	waitGroup.Done() // Indicates that the thread ended
}

func bar() {
	for i := 0; i < limit; i++ {
		fmt.Println("bar: ", i)
	}
}
