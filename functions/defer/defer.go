package main

import (
	"fmt"
)

func main() {
	defer foo() // defers the running of the function until the exit of the main function
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
