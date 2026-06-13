package main

import (
	"fmt"
	"time"
)

func main() {

	// Buffered channel with capacity 1.
	c1 := make(chan string, 1)

	// Start a goroutine that takes 2 seconds
	// before sending a result.
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Wait for whichever happens first:
	//   1. Receive from c1
	//   2. Timeout after 1 second
	select {

	// Runs if c1 receives a value first.
	case res := <-c1:
		fmt.Println(res)

	// Runs if 1 second passes first.
	//
	// time.After() returns a channel that
	// receives a value after the given duration.
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// Another buffered channel.
	c2 := make(chan string, 1)

	// This goroutine also takes 2 seconds.
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	// Wait for whichever happens first:
	//   1. Receive from c2
	//   2. Timeout after 3 seconds
	select {

	// c2 sends after 2 seconds.
	case res := <-c2:
		fmt.Println(res)

	// Timeout would happen after 3 seconds.
	// It loses the race.
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
