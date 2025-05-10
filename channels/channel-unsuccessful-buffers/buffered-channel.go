package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1) // Create a buffered channel with a capacity of 1 so that it can hold one value until it is read by another goroutine

	c <- 42
	c <- 43 // This will block the program because the channel is full and there is no receiver to read the value; and the program will deadlock because there is no goroutine to read the value from the channel

	fmt.Println(<-c)

}
