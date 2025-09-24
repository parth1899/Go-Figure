package main

// Buffered channels

// channel buffering is FIFO, channels in Go are FIFO (first-in, first-out) queues
// Buffered channels accept a limited number of values without a corresponding receiver for those values.

import "fmt"

func main() {
	c := make(chan int, 3) // second argument is the buffer size
	// Sends to a buffered channel blocks only when the buffer is full.
	// Receives blocks when the buffer is empty.

	c <- 1
	c <- 2
	c <- 3

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
