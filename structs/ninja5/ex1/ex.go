package main

import (
	"fmt"
)

type person struct {
	first              string
	last               string
	ice_cream_flavours []string
}

func main() {
	p1 := person{
		first:              "Person",
		last:               "1",
		ice_cream_flavours: []string{"Rocky Road", "Golden Ribbon"},
	}

	p2 := person{
		first:              "Person",
		last:               "2",
		ice_cream_flavours: []string{"Missisippi Mud", "Golden Ribbon"},
	}

	for _, v := range p1.ice_cream_flavours {
		fmt.Printf("%v %v's favourite ice cream flavour is %v\n", p1.first, p1.last, v)
	}

	for _, v := range p2.ice_cream_flavours {
		fmt.Printf("%v %v's favourite ice cream flavour is %v\n", p2.first, p2.last, v)
	}
}
