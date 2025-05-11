package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 33
	}()

	v, ok := <-c
	fmt.Println(v, ok) // 33 true

	close(c)

	v, ok = <-c
	fmt.Println(v, ok) // 0 false
}
