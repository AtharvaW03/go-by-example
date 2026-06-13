package main

import (
	"fmt"
	"time"
)

// Worker receives a channel used to notify
// the caller when the work is finished.
func worker(done chan bool) {

	// Simulate some work.
	fmt.Print("working...")

	// Pretend the work takes 1 second.
	time.Sleep(time.Second)

	fmt.Println("done")

	// Send a value into the channel.
	//
	// This acts as a "finished" signal.
	done <- true
}

func main() {

	// Create a buffered channel that can hold
	// one bool value.
	done := make(chan bool, 1)

	// Start worker in a new goroutine.
	// Main continues immediately.
	go worker(done)

	// Wait for a value from the channel.
	//
	// Main blocks here until worker sends:
	//
	//     done <- true
	//
	// We don't care about the actual value,
	// only that something was received.
	<-done
}
