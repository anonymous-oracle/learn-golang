package main

import (
	"fmt"
)

func main() {
	foo()
	bar("woo")
	s1 := woo("bar")
	fmt.Println(s1)

	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	println(y)

}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello, ", s)
}

func woo(s string) string { // specifying return type of the function after the function parameters
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(fn string, ln string) (string, bool) { // multiple return types
	a := fmt.Sprint(fn, ln, `says, "Hello"`)
	b := false

	return a, b
}
