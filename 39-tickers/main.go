package main

import (
	"fmt"
	"time"
)

func main() {

	// Create a ticker that fires every 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)

	// Used to tell the goroutine to stop.
	done := make(chan bool)

	// Start a background goroutine.
	go func() {

		// Run forever until told to stop.
		for {

			// Wait for either:
			// 1. A stop signal
			// 2. A ticker event
			select {

			// Stop requested.
			case <-done:
				return

			// Ticker fired.
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Let the ticker run for a while.
	time.Sleep(1600 * time.Millisecond)

	// Stop future ticks.
	ticker.Stop()

	// Tell the goroutine to exit.
	done <- true

	fmt.Println("Ticker stopped")
}
