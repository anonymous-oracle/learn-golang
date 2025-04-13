package main

import "fmt"

func main()  {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println("Printing on a single line")
	fmt.Printf("%v\t%v\t%v\n",x,y,z)

	fmt.Println("Printing on multiple lines")
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)

}