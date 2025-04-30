package main

import (
	"fmt"
)

var x int // the scope of x is available for the entire code block

func main()  {
	fmt.Println(x)
	x++
	fmt.Println(x)
}