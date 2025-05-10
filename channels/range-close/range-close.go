package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) // close the channel after sending all values; not closing c will cause deadlock because the receiver will wait for more values to come; closing channels must be done in the same goroutine that sends values to it
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("exiting main...")

}
