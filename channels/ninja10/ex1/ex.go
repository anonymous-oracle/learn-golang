package main

import (
	"fmt"
)

func main() {
	// // Solution 1 - using a buffered channel
	// c := make(chan int, 1)

	// c <- 42

	// fmt.Println(<-c)

	// Solution 2 - using a buffered channel with a for loop
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

}
