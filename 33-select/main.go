package main

import (
	"fmt"
	"time"
)

func main() {

	// Two unbuffered channels.
	c1 := make(chan string)
	c2 := make(chan string)

	// Send a message to c1 after 1 second.
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	// Send a message to c2 after 2 seconds.
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// We expect two messages total.
	for range 2 {

		// Wait until one of the channels is ready.
		select {

		// Runs if c1 receives first.
		case msg1 := <-c1:
			fmt.Println("received", msg1)

		// Runs if c2 receives first.
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
