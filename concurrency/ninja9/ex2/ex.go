package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Printf("My name is %s and I am %d years old.\n", p.name, p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "John",
		age:  30,
	}
	p2 := person{
		name: "Jane",
		age:  25,
	}

	saySomething(&p1)
	saySomething(&p2)
}
