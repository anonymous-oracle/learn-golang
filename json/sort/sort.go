package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println(xi)
	fmt.Println(xs)

	// sorting the slices

	sort.Ints(xi)    // sorts in place
	sort.Strings(xs) // sorts in place

	fmt.Println(xi)
	fmt.Println(xs)

}
