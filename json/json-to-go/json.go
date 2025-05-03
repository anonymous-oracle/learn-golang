package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"` // these are json tags; necessary for unmarshalling
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Sung","Last":"Jinwoo","Age":29},{"First":"Thomas","Last":"Andre","Age":41}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range people {
		fmt.Println(v)
	}

}
