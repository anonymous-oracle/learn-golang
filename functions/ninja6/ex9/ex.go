package main

import (
	"fmt"
)

func main() {
	fmt.Println(calling_func(callback, "Leveling"))
}

func callback(s string) string {
	return fmt.Sprintf("Solo %s", s)
}

func calling_func(f func(s string) string, s string) string {
	return f(s)
}
