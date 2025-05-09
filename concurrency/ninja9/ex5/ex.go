package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var counter int64
	
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			// time.Sleep(time.Microsecond)
			runtime.Gosched() // Yield the processor, allowing other goroutines to run
			fmt.Println("Counter:", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
	println(counter)
}
