package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string // if the data needs to be Marshalled to JSON, ensure that the field names start with Upper case letters
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Sung",
		Last:  "Jinwoo",
		Age:   29,
	}

	p2 := person{
		First: "Thomas",
		Last:  "Andre",
		Age:   41,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}
	//	for _, v := range bs {
	//		fmt.Println(string(v))
	//	}
	fmt.Println(string(bs)) // viewing the json encoded data

}
