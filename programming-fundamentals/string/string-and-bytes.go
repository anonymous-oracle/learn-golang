package main

import (
	"fmt"
)

func main() {
	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i]) /// %#U gives the UTF-8 code point of the byte (or a byte of string)
	}
	fmt.Printf("\n")
	for i := range len(s) {
		fmt.Printf("%#U", s[i]) // %#U gives the UTF-8 code point of the byte (or a byte of string)
	}
	fmt.Printf("\n")

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}

	s = "H"
	fmt.Println(s)

	bs = []byte(s)
	fmt.Println(bs)
	fmt.Printf("%#U\n", bs[0])

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#X\n", n)

}
