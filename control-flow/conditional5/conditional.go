package main

import (
	"fmt"
)

func main() {
	// switch "Bond" {
	n := "Bond"
	switch n {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("bond, james")
	case "bond", "Q", "skyfall":
		fmt.Println("Multi expression matching case")
	default:
		fmt.Println("this is default")
	}
}
