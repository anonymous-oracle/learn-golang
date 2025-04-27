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
		ice_cream_flavours: []string{"Missisippi Mud", "Choco Chip"},
	}

	var person_map = map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range person_map {
		fmt.Println(person_map[k])
		for _, iv := range v.ice_cream_flavours {
			fmt.Printf("%v %v's favourite ice cream flavour is %v\n", v.first, v.last, iv)
		}
	}
}
