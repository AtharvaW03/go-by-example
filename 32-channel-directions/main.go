package main

import "fmt"

// pings is a send-only channel.
//
// Read:
// string -> pings
//
// This function can send strings into pings,
// but cannot receive from it.
func ping(pings chan<- string, msg string) {

	// Send msg into the channel.
	pings <- msg
}

// pings is a receive-only channel.
// Read:
// pings -> string
//
// pongs is a send-only channel.
// Read:
// string -> pongs
func pong(pings <-chan string, pongs chan<- string) {

	// Receive a string from pings.
	msg := <-pings

	// Send that same string into pongs.
	pongs <- msg
}

func main() {

	// Buffered channels with capacity 1.
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// Put a message into pings.
	ping(pings, "passed message")

	// Move the message from pings to pongs.
	pong(pings, pongs)

	// Receive from pongs and print.
	fmt.Println(<-pongs)
}
