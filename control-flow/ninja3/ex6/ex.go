package main

import (
	"fmt"
)

func main() {
	x := 1

	if x > 1 {
		fmt.Println("x is greater than 1")
	} else if x < 1 {
		fmt.Println("x is lesser than 1")
	} else {
		fmt.Println("x is equal to 1")
	}
}
