package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2) // Create a buffered channel with a capacity of 1 so that it can hold one value until it is read by another goroutine

	c <- 42
	c <- 43

	fmt.Println(<-c) // pulls of the first value from the channel
	// fmt.Println(<-c) // Read the value from the channel
}
