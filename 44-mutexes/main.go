package main

import (
	"fmt"
	"sync"
)

// Container holds shared data and a mutex to protect it.
type Container struct {

	// Mutex used to ensure only one goroutine
	// modifies the map at a time.
	mu sync.Mutex

	// Shared counters.
	counters map[string]int
}

// Increment a counter safely.
func (c *Container) inc(name string) {

	// Acquire lock.
	// If another goroutine already holds it,
	// wait until it is released.
	c.mu.Lock()

	// Always release the lock when the
	// function returns.
	defer c.mu.Unlock()

	// Safe map update.
	c.counters[name]++
}

func main() {

	// Create a container with two counters.
	c := Container{
		counters: map[string]int{
			"a": 0,
			"b": 0,
		},
	}

	// Used to wait for all goroutines to finish.
	var wg sync.WaitGroup

	// Anonymous helper function that increments a counter
	// n times.
	doIncrement := func(name string, n int) {

		for range n {
			c.inc(name)
		}
	}

	// Increment counter "a" 10,000 times.
	wg.Go(func() {
		doIncrement("a", 10000)
	})

	// Increment counter "a" another 10,000 times.
	wg.Go(func() {
		doIncrement("a", 10000)
	})

	// Increment counter "b" 10,000 times.
	wg.Go(func() {
		doIncrement("b", 10000)
	})

	// Wait until all goroutines finish.
	wg.Wait()

	// Print final counter values.
	fmt.Println(c.counters)
}
