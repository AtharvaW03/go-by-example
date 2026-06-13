package main

import "fmt"

func main() {

	// Buffered channel that can hold up to 5 jobs.
	jobs := make(chan int, 5)

	// Used to signal when the worker has finished.
	done := make(chan bool)

	// Start worker goroutine.
	go func() {

		// Keep receiving jobs until the channel is closed.
		for {

			// Receive a job.
			//
			// j    = job value
			// more = true  -> channel still open
			//        false -> channel closed and empty
			j, more := <-jobs

			if more {

				// Process the job.
				fmt.Println("received job", j)

			} else {

				// No more jobs will ever arrive.
				fmt.Println("received all jobs")

				// Notify main that we're finished.
				done <- true

				// Exit the goroutine.
				return
			}
		}
	}()

	// Send 3 jobs into the jobs channel.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	// Close the jobs channel.
	//
	// This means:
	// "No more jobs will be sent."
	//
	// Existing jobs can still be received.
	close(jobs)

	fmt.Println("sent all jobs")

	// Wait for the worker to finish.
	//
	// Main pauses here until the worker does:
	//
	//     done <- true
	//
	<-done

	// Try to receive from jobs again.
	//
	// Since the channel is closed and empty:
	//
	// ok = false
	//
	// This is how Go tells us there are
	// no more values available.
	_, ok := <-jobs

	fmt.Println("received more jobs:", ok)
}
