package main

import "fmt"

var a int

type hotdog int // defining a custom type that will have an underlying basic data type
var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// a = b // even if the underlying type of b is int, assigning the value of b to a is invalid

	// ENTER typecasting or in golang, typecasting is specifically called conversion

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
