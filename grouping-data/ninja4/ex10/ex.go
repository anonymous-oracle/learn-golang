package main

import (
	"fmt"
)

func main() {
	var m = map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	for key, val := range m {
		fmt.Println("The list of values for the key ", key)
		for i, v := range val {
			fmt.Println(i, v)
		}
	}

	m["jinwoo_sung"] = []string{"Leveling up", "Shadows", "Daggers", "Rizz"}
	fmt.Printf("\n\n")
	for key, val := range m {
		fmt.Println("The list of values for the key ", key)
		for i, v := range val {
			fmt.Println(i, v)
		}
	}

	delete(m, "no_dr")
	fmt.Printf("\n\n")
	for key, val := range m {
		fmt.Println("The list of values for the key ", key)
		for i, v := range val {
			fmt.Println(i, v)
		}
	}
}
