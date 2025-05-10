package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// send
	go foo(c)

	// receive
	// go bar(c) // a separate goroutine to receive may not actually work if the goroutine has to wait for reading from the channel and the main() function can likely exit well before that happens
	bar(c) // this will hold the execution until the value is received from the channel

	fmt.Println("exiting...")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
