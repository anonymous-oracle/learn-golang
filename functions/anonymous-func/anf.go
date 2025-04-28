package main

import (
	"fmt"
)

func main() {
	foo()
	func() {
		fmt.Println("Anonymous func ran")
	}() // adding parantheses after declaring the anonymous function
	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42) // this is with passing arguments to the function call immediately after declaration with parameters
}

func foo() {
	fmt.Println("foo ran")
}
