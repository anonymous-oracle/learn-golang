package main

import (
	"fmt"
)

func main() {
	m := map[string]int{ // the key type is enclosed within [] and the value is outside
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])
	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok) // unlike python, if the element does not exist and [] is used, a key error is not thrown

	if v, ok := m["Barnabas"]; ok {
		fmt.Println("This is the if print for Barnabas", v)
	}

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("This is the if print for Miss Moneypenny", v)
	}

	m["todd"] = 33
	fmt.Println("\nPrinting out the key value pairs of the map...")
	for k, v := range m {
		fmt.Printf("%v -> %v\n", k, v)
	}

	// deleting a key from the map is easy
	delete(m, "todd")
	fmt.Println(m)

	delete(m, "Ian Fleming")
	fmt.Println(m)

	// to make sure something is deleted
	if _, ok := m["todd"]; ok {
		fmt.Println("Value -> ", "todd", "still present")
		delete(m, "todd")
	}
}
