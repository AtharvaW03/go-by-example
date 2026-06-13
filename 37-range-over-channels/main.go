package main

import "fmt"

func main() {

	// Create a buffered channel with capacity 2.
	queue := make(chan string, 2)

	// Send two values into the channel.
	queue <- "one"
	queue <- "two"

	// Close the channel.
	// No more values can be sent after this.
	close(queue)

	// Receive values from the channel until it is
	// closed and all buffered values are consumed.
	for elem := range queue {

		// Print each received value.
		fmt.Println(elem)
	}
}
