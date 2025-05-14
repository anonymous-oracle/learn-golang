package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("Hello, World!")

	io.Copy(f, r)

}
