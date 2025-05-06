package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	fmt.Println("done")
}

func foo() {
	fmt.Println("printing something")
	wg.Done()
}

func bar() {
	fmt.Println("printing something else")
	wg.Done()
}
