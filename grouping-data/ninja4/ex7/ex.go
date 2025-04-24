package main

import (
	"fmt"
)

func main() {
	r1 := []string{"James", "Bond", "Shaken, not stirred"}
	r2 := []string{"Miss", "Moneypenny", "Hello James"}

	mslice := [][]string{r1, r2}

	for _, v := range mslice {
		for _, rv := range v {
			fmt.Println(rv)
		}
	}

}
