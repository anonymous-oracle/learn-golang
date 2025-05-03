package main

import (
	"fmt"
)

func main() {
	returned_func := func_return()
	fmt.Print(returned_func())
}

func func_return() func() string {
	return func() string {
		return fmt.Sprintln("Printing from returned function")
	}
}
