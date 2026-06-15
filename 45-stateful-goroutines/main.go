package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Read request.
//
// key  = which key to read
// resp = channel used to send the answer back
type readOp struct {
	key  int
	resp chan int
}

// Write request.
//
// key  = which key to update
// val  = new value
// resp = channel used to confirm completion
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// Counters tracking how many read and write
	// operations were completed.
	var readOps uint64
	var writeOps uint64

	// Channels used to send requests to the
	// state manager goroutine.
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// State manager goroutine.
	//
	// This goroutine is the ONLY code that
	// directly accesses the map.
	go func() {

		// Shared state.
		var state = make(map[int]int)

		// Process requests forever.
		for {
			select {

			// Handle read request.
			case read := <-reads:

				// Look up value and send it back
				// through the response channel.
				read.resp <- state[read.key]

			// Handle write request.
			case write := <-writes:

				// Update state.
				state[write.key] = write.val

				// Tell caller we're done.
				write.resp <- true
			}
		}
	}()

	// Start 100 reader goroutines.
	for range 100 {

		go func() {

			// Read forever.
			for {

				// Create a read request.
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}

				// Send request to state manager.
				reads <- read

				// Wait for response.
				<-read.resp

				// Increment read counter safely.
				atomic.AddUint64(&readOps, 1)

				// Small pause.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Start 10 writer goroutines.
	for range 10 {

		go func() {

			// Write forever.
			for {

				// Create a write request.
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}

				// Send request to state manager.
				writes <- write

				// Wait for confirmation.
				<-write.resp

				// Increment write counter safely.
				atomic.AddUint64(&writeOps, 1)

				// Small pause.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let everything run for 1 second.
	time.Sleep(time.Second)

	// Safely read final operation counts.
	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)

	// Print results.
	fmt.Println("readOps:", readOpsFinal)
	fmt.Println("writeOps:", writeOpsFinal)
}
