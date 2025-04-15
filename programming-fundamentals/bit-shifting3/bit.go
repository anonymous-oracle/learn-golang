package main

import (
	"fmt"
)

const (
	_ = iota
	// kb = 1024
	kb = 1 << (iota * 10) // a kilo byte is 1 shifted 10 times in the binary number system; iota value is 1
	mb = 1 << (iota * 10) // a mega byte is 1 kilo byte shifted 10 times in the binary number system; iota value is 2
	gb = 1 << (iota * 10) // a giga byte is 1 mega byte shifted 10 times in the binary number system; iota value is 3
)

func main() {

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
