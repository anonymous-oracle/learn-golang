package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("printing from an anonymous function")
	}()
}
