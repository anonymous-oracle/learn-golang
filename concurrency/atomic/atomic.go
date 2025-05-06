package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"golang.org/x/telemetry/counter"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	var counter int64

	const gs = 100
	var wg sync.WaitGroup // initialising the wait group
	wg.Add(gs)            // adding gs number of wait groups

	for i := 0; i < gs; i++ {
		go func() {

			counter++
			runtime.Gosched() // yield the thread so that the execution of another go routine can start/continue

			wg.Done() // calling done for each function literal execution present in the wait group routine

		}()

		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait() // waits for all routines to end before exiting the main function
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	fmt.Println("count:", counter)

}
