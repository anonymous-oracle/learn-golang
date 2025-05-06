package main

import (
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	counter := 0

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			time.Sleep(time.Microsecond)
			runtime.Gosched() // Yield the processor, allowing other goroutines to run
			v++
			counter = v
			wg.Done()
		}()
	}

	wg.Wait()
	println(counter)
}
