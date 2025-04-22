package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 101}
	for _, v := range slice {
		fmt.Println(v)
	}

	s1 := append(slice[:4], slice[6:]...)

	fmt.Println(s1)
}
