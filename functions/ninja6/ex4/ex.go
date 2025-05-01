package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hi, I'm %s and I am %d years old\n", p.first, p.age)
}

func main() {
	var p person
	p = person{
		first: "Sung",
		last:  "Jinwoo",
		age:   29,
	}
	p.speak()
}
