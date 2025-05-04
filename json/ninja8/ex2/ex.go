package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	j := `[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]`

	var people []person
	err := json.Unmarshal([]byte(j), &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

}
