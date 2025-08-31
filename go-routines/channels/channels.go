package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values into another goroutine
// <- is the channel operator

func sum(n []int, c chan int) {
	sum := 0

	for _, v := range n {
		sum += v
	}

	c <- sum // the channel c will block this "sends (c <-)" until there is a corresponding "receives (<- c)" ready to receive the sent value.

	// The sending goroutine blocks until another goroutine executes <-c to receive that value.
}

func main() {

	s := []int{1, 2, 3, 4, 5, 6}

	c := make(chan int) // Create a new channel with make(chan val-type). Channels are typed by the values they convey.

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

}

// By default, sends and receives block until both the sender and receiver are ready.
// This property allows goroutines to synchronize without explicit locks or condition variables.

// In Go, channels force synchronization: a send and a receive must meet.

// When you do c <- sum (send a value into a channel):
// The sending goroutine will pause (block) until some other goroutine is ready to receive from that channel (<-c).
// Similarly, when you do x := <-c (receive a value from a channel):
// The receiving goroutine will pause (block) until some other goroutine sends a value into the channel.

// In Go, we don’t usually need locks or condition variables Because of the blocking property:
// A goroutine that sends knows that the other goroutine has received when the send returns.
// A goroutine that receives knows that some goroutine has sent when the receive retur
