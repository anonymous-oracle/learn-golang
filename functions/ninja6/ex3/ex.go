package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo runs second, at the time of exit of the main function as it is deferred")
}

func bar() {
	fmt.Println("bar runs first as it is not deferred")
}
