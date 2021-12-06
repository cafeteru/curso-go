package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("GO routines\t", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var mutex sync.Mutex
	var waitGroup sync.WaitGroup
	waitGroup.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mutex.Lock() // Lock variable
			v := counter
			v++
			runtime.Gosched()
			counter = v
			fmt.Println("Counter: ", counter)
			mutex.Unlock() // Unlock variable
			waitGroup.Done()
		}()
		fmt.Println("GO routines\t", runtime.NumGoroutine())
	}
	waitGroup.Wait()
	fmt.Println("Counter: ", counter)
}
