package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1) // Create a buffered channel with a capacity of 1 so that it can hold one value until it is read by another goroutine

	c <- 42

	fmt.Println(<-c)

}
