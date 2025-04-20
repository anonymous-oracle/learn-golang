package main

import (
	"fmt"
)

func main() {
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))

	fmt.Println(cap(x)) // max values that can be stored in the current memory allocation of the slice
	x[0] = 42
	x[9] = 999

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// x[10] = 3423 // this will throw an error
	x = append(x, 3423)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3424)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3425)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	
	

}
