package main

import (
	"fmt"
)

func main() {
	switch {
	case true:
		fmt.Println("Since the expression is empty, the execution defaults to the true case")
	case false:
		fmt.Println("Will not execute unless 'fallthrough' is used")
	default:
		fmt.Println("Will not execute since there is a true case")

	}
}
