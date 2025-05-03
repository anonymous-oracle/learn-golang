package main

import (
	"fmt"
)

func main() {
	x := 0
	foo(x)
	println(x) // x is still 0 because everything in go is pass by value
}

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}
