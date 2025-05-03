package main

import (
	"fmt"
)

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a
	fmt.Println(b)
	c := *b // de-referencing the address/pointer
	fmt.Println(c)

	fmt.Println(*&a) // de-referencing the address directly

	*b = 43
	fmt.Println(a)

}
