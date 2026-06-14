package main

import (
	"fmt"
	"time"
)

// Worker receives jobs and sends results.
func worker(id int, jobs <-chan int, results chan<- int) {

	// Keep receiving jobs until the jobs
	// channel is closed.
	for j := range jobs {

		fmt.Println("worker", id, "started job", j)

		// Simulate work.
		time.Sleep(time.Second)

		fmt.Println("worker", id, "finished job", j)

		// Send result back.
		results <- j * 2
	}
}

func main() {

	const numJobs = 5

	// Job queue.
	jobs := make(chan int, numJobs)

	// Result queue.
	results := make(chan int, numJobs)

	// Start 3 worker goroutines.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs to workers.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// No more jobs will be sent.
	close(jobs)

	// Wait for all results.
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
