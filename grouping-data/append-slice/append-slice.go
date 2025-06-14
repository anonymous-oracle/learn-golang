package main

import (
	"fmt"
)

func main() {
	x := []int{4, 5, 6, 8, 9}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)

	fmt.Println(x)
	y := []int{234, 456, 678, 987}
	x = append(x, y...) // unpacks like python *args
	fmt.Println(x)
}
