package main

import (
	"fmt"
)

func main() {
	c := gen() // the go routine feeding to the channel will be blocked until the channel is read

	func() {
		for v := range c { // reads off of the channel until the channel is closed
			fmt.Println(v)
		}
	}()

	fmt.Println("Exiting main...")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}() // this routine will execute separately and return the active channel in the main execution routine
	return c
}
