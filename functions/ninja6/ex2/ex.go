package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum([]int{1, 2, 6, 8, 9, 6, 9, 79, 3, 3, 4, 9, 47, 9, 4, 6, 45}...))
}

func sum(xsl ...int) int {
	total_sum := 0
	for _, v := range xsl {
		total_sum += v
	}
	return total_sum
}
