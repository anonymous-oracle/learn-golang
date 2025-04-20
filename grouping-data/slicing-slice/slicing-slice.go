package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 6, 8, 9}
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[3:])
	fmt.Println(x[1:3])

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}
