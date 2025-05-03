package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func changeMe(p *person) {
	(*p).name = "I HAVE, NO NAME!!"
	(*p).age = 45649
}

func main() {
	p := person{
		name: "Zangetsu",
		age:  16,
	}

	fmt.Println(p)
	changeMe(&p)

	fmt.Println(p)
}
