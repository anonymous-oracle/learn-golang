package main

import (
	"fmt"
)

func main() {
	x := 2
	fmt.Printf("%d\t\t%b\n", x, x)
	x = 4
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)
}
