package main

import (
	"fmt"
)

func main() {
	// cs := make(chan<- int) // the problem here is that the channel is a send-only channel
	cs := make(chan int)

	go func() {
		cs <- 42
	}()

	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

}
