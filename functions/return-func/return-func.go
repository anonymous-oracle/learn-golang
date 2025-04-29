package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())

	x := bar()

	fmt.Printf("%T\n", bar())

	fmt.Println(x())
	fmt.Println(bar()())

}

func foo() string {
	// s := "Hello world"
	// return s
	return "Hello world"
}

// bar() has a return type of an anonymous function which inturn returns a type of int; within the body, an anonymous function with return type of int is being returned
func bar() func() int {
	return func() int {
		return 451
	}
}
