package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	counter := 0

	const gs = 100
	var wg sync.WaitGroup // initialising the wait group
	wg.Add(gs)            // adding gs number of wait groups

	for i := 0; i < gs; i++ {
		go func() {
			v := counter

			time.Sleep(time.Microsecond) // sleeps for a second and suspends the execution for the same duration
			runtime.Gosched()            // yield the thread so that the execution of another go routine can start/continue

			v++
			counter = v

			wg.Done() // calling done for each function literal execution present in the wait group routine

			fmt.Println("Goroutines:", runtime.NumGoroutine())
		}()
	}
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	fmt.Println("count:", counter)

	wg.Wait() // waits for all routines to end before exiting the main function
}
