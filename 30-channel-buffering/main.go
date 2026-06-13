package main

import "fmt"

func main() {

	// Create a buffered channel.
	//
	// Type: string
	// Capacity: 2
	//
	// The channel can hold up to 2 strings
	// before a receiver is required.
	messages := make(chan string, 2)

	// Put first value into the buffer.
	messages <- "buffered"

	// Put second value into the buffer.
	messages <- "channel"

	// Buffer is now full.
	//
	// A third send here would block:
	//
	// messages <- "oops"

	// Receive the first value.
	// Channels are FIFO (First In, First Out).
	fmt.Println(<-messages)

	// Receive the second value.
	fmt.Println(<-messages)
}
