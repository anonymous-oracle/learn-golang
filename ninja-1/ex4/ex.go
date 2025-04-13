package main

import "fmt"

type ownint int

var x ownint

func main() {
	fmt.Printf("Value of x is %v\n", x)
	fmt.Printf("Type of x is %T\n", x)
	x = 42
	fmt.Printf("Value of x is %v\n", x)
}
