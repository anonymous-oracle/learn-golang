package main

import (
	"fmt"
)

func main() {
	// a slice uses something called a composite literal
	// x := type{values}
	x := []int{4, 5, 6, 7, 31}
	fmt.Println(x)
}
