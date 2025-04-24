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

	for _, val := range m {
		for i, v := range val {
			fmt.Println(i, v)
		}
	}

}
