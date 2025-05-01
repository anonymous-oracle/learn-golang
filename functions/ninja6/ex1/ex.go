package main

import (
	"fmt"
)

func main() {
	fmt.Println("foo results:", foo())
	b1, b2 := bar()
	fmt.Println("bar results:", b1, b2)
}

func foo() int {
	return 8
}

func bar() (int, string) {
	return 8, "Eight"
}
