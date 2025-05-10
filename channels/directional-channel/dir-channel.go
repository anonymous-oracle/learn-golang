package main

import (
	"fmt"
)

func main() {
	c := make(chan<- int, 2) // creating a send-only directional channel
	// c := make(<- chan int, 2) // creating a receive-only directional channel

	c <- 42
	c <- 43

	fmt.Println(<-c) // This line will cause a compilation error because c is a send-only channel
	fmt.Println(<-c) // This line will also cause a compilation error because c is a send-only channel
	fmt.Println("-----")
	fmt.Printf("%T\n", c)

	// NOTE: a general/bidirectional channel can be assigned a directional channel, but a receive-only channel cannot be assigned a send-only channel and vice versa. A directional channel can be assigned to a directional channel
}
