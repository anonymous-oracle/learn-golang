package main

import (
	"fmt"
)

func main() {
	// a slice uses something called a composite literal
	// x := type{values}
	x := []int{4, 5, 6, 7, 31}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i, v := range x { // returns an index and the corresponding value
		fmt.Println(i, v)
	}

	xi := []int{4, 5, 7, 8, 9, 42}
	for i, v := range xi {
		fmt.Println(i, v)
	}
}
