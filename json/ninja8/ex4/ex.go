package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 6, 8, 9, 4, 2, 55, 6, 6, 87, 74, 4, 4, 456, 9, 94, 9, 499, 94, 94, 94, 8, 49, 48, 49, 49, 8489, 498, 4}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry"}

	fmt.Println(xi)

	sort.Ints(xi)

	fmt.Println(xi)

	fmt.Println(xs)

	sort.Strings(xs)

	fmt.Println(xs)
}
