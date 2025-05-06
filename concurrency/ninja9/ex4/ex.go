package main

import (
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter := 0
	var mu sync.Mutex
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			v := counter
			// time.Sleep(time.Microsecond)
			runtime.Gosched() // Yield the processor, allowing other goroutines to run
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	println(counter)
}
