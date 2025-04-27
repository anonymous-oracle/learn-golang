package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle   vehicle
	fourWheel bool
}

type sedan struct {
	vehicle vehicle
	luxury  bool
}

func main() {
	var truck = truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	var sedan = sedan{
		vehicle: vehicle{
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
