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
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("exiting from main... ")
}
