package main

import (
	"fmt"
)

func main() {
	x := 47
	for {
		x++
		if x > 122 {
			break
		}
		fmt.Printf("%v\t%#x\t%#U\n", x, x, x) // displays the original value, hexadecimal and the unicode point with the corresponding rune
	}
	fmt.Println("done.")
}
