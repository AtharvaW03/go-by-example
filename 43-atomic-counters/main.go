package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Thread-safe counter.
	var ops atomic.Uint64

	// Wait for all goroutines.
	var wg sync.WaitGroup

	// Start 50 goroutines.
	for range 50 {

		wg.Go(func() {

			// Increment counter 1000 times.
			for range 1000 {

				// Atomic increment.
				//
				// Safe even when many goroutines
				// execute this simultaneously.
				fmt.Println(ops.Add(1))
			}
		})
	}

	// Wait until all goroutines finish.
	wg.Wait()

	// Safely read current value.
	fmt.Println("ops:", ops.Load())
}
