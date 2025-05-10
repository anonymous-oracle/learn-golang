package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42 // this channel will not block because it is buffered and executes in a separate goroutine
	}()

	fmt.Println(<-c) // this will block until the value is received from the channel
}
