package main

import (
	"fmt"
)

// const a = 42
// const b = 42.78
// const c = "Solo Leveling"

// const (
// 	a = 42
// 	b = 42.78
// 	c = "Solo Leveling"
// )

const (
	a int     = 42
	b float64 = 42.78
	c string  = "Solo Leveling"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}
