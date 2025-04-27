package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}
type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}
	fmt.Println(p1.age, p1.first, p1.last)
	fmt.Println(p2.age, p2.first, p2.last)
	fmt.Println(sa1.age, sa1.first, sa1.last, sa1.ltk)

}
