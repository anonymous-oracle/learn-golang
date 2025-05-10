package main

func main() {
	c := make(chan int)

	c <- 42 // put 42 onto c

	println(<-c) // read off of c

	// NOTE: This code will not run because the channel is unbuffered and there is no goroutine to read from it, when reading and writing into the channel, the execution is blocked. There is no goroutine to read from the channel, so the program will deadlock.
}
