package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	foo(&x)
	fmt.Println("x after", &x) // x is still 0 because everything in go is pass by value
	fmt.Println("x after", x)
}

func foo(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
