package main

import (
	"fmt"
)

// const (
// 	a = iota // the iota key word when initialised to a constant, it will store int values starting from 0 in the first constant and then stores the incremented value in all successive const declarations that have no value assigned or is left blank
// 	b = iota
// 	c = iota
// )

const (
	a = iota // the iota key word when initialised to a constant, it will store int values starting from 0 in the first constant and then stores the incremented value in all successive const declarations that have no value assigned or is left blank
	b
	c
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
