package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code }
func (s secretAgent) speak() { // this is a format of attaching a method to a struct by using the receiver clause
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

type human interface {
	speak() // any type that has this method is also of type human
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar as a person", h.(person).first) // unlike conversion which is similar to type casting in other languages, this here demonstrates assertion to an underlying type
	case secretAgent:
		fmt.Println("I was passed into bar as a secret agent", h.(secretAgent).first)
	}
	// human.speak()
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()

	// a value can be of more than one type
	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(p1)
	p1.speak()
	bar(sa1)
	bar(p1)
}
