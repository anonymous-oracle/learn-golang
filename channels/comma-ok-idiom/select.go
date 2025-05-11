package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)

	fmt.Println("exiting main...")

}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}

	}
	close(q) // closing the quit channel to signal the receiver to stop
}

func receive(e, o <-chan int, q <-chan bool) {
	for { // doing until the channels close
		select {
		case v := <-e:
			fmt.Println("from the eve channel:", v)

		case v := <-o:
			fmt.Println("from the odd channel:", v)

		case i, ok := <-q:
			if !ok { // if the channel is closed returns false
				fmt.Println("from comma ok bit", i, ok)
				return
			} else {
				fmt.Println("from comma ok bit", i, ok)
			}
		}

	}
}
