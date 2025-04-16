package main

import (
	"fmt"
)
const (
	year_off = 2020
	_ = iota
	y1 = year_off + iota
	y2 = year_off + iota
	y3 = year_off + iota
	y4 = year_off + iota

)
func main()  {
	fmt.Println(y1, y2, y3, y4)
}