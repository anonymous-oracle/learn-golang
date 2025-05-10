package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)

	fmt.Println("exiting main...")

}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}

	}
	// close(e)
	// close(o)

	q <- 0
	// close(q)
}

func receive(e, o, q <-chan int) {
	for { // doing until the channels close
		select {
		case v := <-e:
			fmt.Println("from the eve channel:", v)

		case v := <-o:
			fmt.Println("from the odd channel:", v)

		case v := <-q:
			fmt.Println("from the quit channel:", v)
			return // returning since a perpetual loop is being used to read from channels; because of this closing the channels is not necessary in this case
		}

	}
}
