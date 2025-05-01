package main

import (
	"fmt"
)

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a()) // prints 4 since the a now has an initialised variable x with 0
	fmt.Println(b()) // but since incrementor is called again, it is creating a new variable in memory with x in b's scope
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
