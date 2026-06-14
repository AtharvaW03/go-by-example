package main

import (
	"fmt"
	"time"
)

func main() {

	// Create a timer that fires after 2 seconds.
	timer1 := time.NewTimer(2 * time.Second)

	// Wait until the timer fires.
	<-timer1.C

	fmt.Println("Timer 1 fired")

	// Create another timer that would fire after 1 second.
	timer2 := time.NewTimer(time.Second)

	// Goroutine waiting for timer2.
	go func() {

		// Wait for timer2 to fire.
		<-timer2.C

		fmt.Println("Timer 2 fired")
	}()

	// Attempt to stop timer2 before it fires.
	stop2 := timer2.Stop()

	// Stop() returns true if the timer was
	// successfully cancelled.
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Wait long enough that timer2 would have fired.
	time.Sleep(2 * time.Second)
}
