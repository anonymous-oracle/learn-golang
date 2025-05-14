package main

import (
	"errors"
	"log"
)

var ErrNorgateMath = errors.New("math: square root of negative number")

func main() {
	log.Printf("%T\n", ErrNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}
func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return 42, nil
}
