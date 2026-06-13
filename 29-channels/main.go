package main

import "fmt"

func main() {

	// Create a channel that can carry string values.
	messages := make(chan string)

	// Start a goroutine.
	go func() {

		// Send "ping" into the channel.
		// The goroutine will wait here until
		// someone receives the value.
		messages <- "ping"
	}()

	// Receive a value from the channel.
	// Main waits here until a value arrives.
	msg := <-messages

	fmt.Println(msg)
}
