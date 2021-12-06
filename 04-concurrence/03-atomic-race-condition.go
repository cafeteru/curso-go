package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUS\t\t", runtime.NumCPU())
	fmt.Println("GO routines\t", runtime.NumGoroutine())
	var counter int64
	const gs = 100
	var waitGroup sync.WaitGroup
	waitGroup.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter: ", counter)
			runtime.Gosched()
			waitGroup.Done()
		}()
		fmt.Println("GO routines\t", runtime.NumGoroutine())
	}
	waitGroup.Wait()
	fmt.Println("Counter: ", counter)
}
