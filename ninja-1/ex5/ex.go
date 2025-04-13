package main

import "fmt"

type ownint int

var x ownint
var y int

func main() {
	fmt.Printf("Value of x is %v\n", x)
	fmt.Printf("Type of x is %T\n", x)
	x = 42
	fmt.Printf("Value of x is %v\n", x)
	y = int(x)
	fmt.Printf("Value of y is %v\n", y)
	fmt.Printf("Type of y is %T\n", y)
}
