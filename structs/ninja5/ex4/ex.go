package main

import (
	"fmt"
)

type truck struct {
	vehicle struct {
		doors int
		color string
	}
	fourWheel bool
}

type sedan struct {
	vehicle struct {
		doors int
		color string
	}
	luxury bool
}

func main() {
	var truck = truck{
		vehicle: struct {
			doors int
			color string
		}{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	var sedan = sedan{
		vehicle: struct {
			doors int
			color string
		}{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}

	fmt.Println(truck)
	fmt.Println(sedan)

	fmt.Println(truck.vehicle)
	fmt.Println(sedan.vehicle)

}
