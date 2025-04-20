package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("this will not print")
	case (2 == 4):
		fmt.Println("this will not print - 2")
	case 3 == 3:
		fmt.Println("this prints")
		fallthrough // unless fallthrough is not specified in the first case that can evaluate to be true, the switch block will exit even if another case could evaluate to true; but if specified all cases after this will be executed even if the cases are false
	case (4 == 4):
		fmt.Println("this is true, but will it print? Sure, just make sure fallthrough is added in the previous true case")
	case 7 == 9:
		fmt.Println("the case is false, but will print because of fallthrough")
		fallthrough
	case 11 == 14:
		fmt.Println("the case is false, but will print because of fallthrough - 2")
		fallthrough
	case 15 == 15:
		fmt.Println("true 15")
	default:
		fmt.Println("this is the default case") // prints if nothing evaluates to true
	}
}
