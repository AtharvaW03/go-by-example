package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulates some work.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// WaitGroup tracks active goroutines.
	var wg sync.WaitGroup

	// Start 5 workers.
	for i := 1; i <= 5; i++ {

		// Start a goroutine and automatically
		// register it with the WaitGroup.
		wg.Go(func() {
			worker(i)
		})
	}

	// Block until all workers finish.
	wg.Wait()
}
