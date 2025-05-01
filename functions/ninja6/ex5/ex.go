package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s square) area() float64 {
	return math.Pow(s.side, 2)
}

type shape interface {
	area() float64 // adding the specific return type will clarify to the compiler about what return type to enforce
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{
		side: 5,
	}
	ci := circle{
		radius: 5,
	}

	info(sq)
	info(ci)

}
