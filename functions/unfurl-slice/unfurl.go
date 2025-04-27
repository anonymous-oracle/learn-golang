package main

import (
	"fmt"
)

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum("String argument", xi...)
	fmt.Println("The total is", x)
}

func sum(s string, x ...int) int { // variadic parameters have to be passed in the end
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding,", v, "to the total which is now")
	}
	return sum
}
