package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i * i
	}

	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", arr)
}
