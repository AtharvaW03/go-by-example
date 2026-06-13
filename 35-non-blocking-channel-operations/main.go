package main

import "fmt"

func main() {

	// Unbuffered channels.
	//
	// They start empty and have no storage.
	// A send requires a receiver to be ready.
	messages := make(chan string)
	signals := make(chan bool)

	// Non-blocking receive.
	//
	// Try to receive from messages.
	// If no message is available immediately,
	// run default instead.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)

	default:
		fmt.Println("no message received")
	}

	msg := "hi"

	// Non-blocking send.
	//
	// Try to send "hi" into messages.
	// Because messages is unbuffered and nobody
	// is receiving, this send cannot proceed.
	//
	// So default runs immediately.
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)

	default:
		fmt.Println("no message sent")
	}

	// Non-blocking select across multiple channels.
	//
	// Try:
	//   receive from messages
	//   receive from signals
	//
	// If neither is ready, run default.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)

	case sig := <-signals:
		fmt.Println("received signal", sig)

	default:
		fmt.Println("no activity")
	}
}
