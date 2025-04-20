package main

import (
	"fmt"
)

func main() {
	ascii_start := 65 // always start from ascii value since it is based on decimals and not hexadecimals
	for x := ascii_start; x < ascii_start+26; x++ {
		fmt.Println(x)
		for range 3 {
			fmt.Printf("\t%#U\n", rune(x))
		}
	}

}
