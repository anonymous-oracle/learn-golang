package main

import (
	"fmt"
	"runtime"
)

func main()  {
	fmt.Println("OS Arch:", runtime.GOARCH)
	fmt.Println("OS:", runtime.GOOS)
} 