package main

import (
	"fmt"
)

func main() {
	s := "This string will have an address"
	fmt.Println(&s)
}
