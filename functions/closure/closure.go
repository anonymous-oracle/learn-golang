package main

import (
	"fmt"
)

var x int // the scope of x is available for the entire code block

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)

	{ // this block ensures the closure of y's scope
		y := 42
		fmt.Println(y) // y has scope limited to this code block
	}
	// fmt.Println(y) // y is out of scope here
}

func foo() {
	fmt.Println("calling from foo")
	x++
}
