package main

import (
	"fmt"
)

func main() {
	for n := 11; n < 100; n++ {
		fmt.Printf("%v %% 4 = %v\n", n, n%4) // use double % to mean only % in a string
	}

}
