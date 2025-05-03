package main

import (
	"fmt"
)

func main() {
	ap1 := slice_appender()
	ap2 := slice_appender()

	fmt.Println(ap1())
	fmt.Println(ap1())
	fmt.Println(ap1())
	fmt.Println(ap1())

	fmt.Println(ap2())
	fmt.Println(ap2())
	fmt.Println(ap2())
	fmt.Println(ap2())

}

func slice_appender() func() []int {
	var x int
	sl := []int{}

	return func() []int {
		x++
		sl = append(sl, x)
		return sl
	}
}
