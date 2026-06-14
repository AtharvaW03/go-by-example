package main

import (
	"fmt"
	"time"
)

func main() {

	// -----------------------------
	// Basic rate limiter
	// -----------------------------

	// Queue of incoming requests.
	requests := make(chan int, 5)

	// Add 5 requests.
	for i := 1; i <= 5; i++ {
		requests <- i
	}

	// No more requests will be sent.
	close(requests)

	// Ticker channel.
	//
	// Produces one value every 200ms.
	limiter := time.Tick(200 * time.Millisecond)

	// Process requests.
	for req := range requests {

		// Wait for the next tick before
		// processing a request.
		<-limiter

		fmt.Println("request", req, time.Now())
	}

	// -----------------------------
	// Bursty rate limiter
	// -----------------------------

	// Token bucket with capacity 3.
	//
	// Think:
	//
	// [T][T][T]
	//
	// Each token allows one request.
	burstyLimiter := make(chan time.Time, 3)

	// Pre-fill bucket with 3 tokens.
	//
	// First 3 requests can proceed immediately.
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// Refill the bucket.
	//
	// Every 200ms, add one token.
	go func() {

		for t := range time.Tick(200 * time.Millisecond) {

			// Add token back into bucket.
			burstyLimiter <- t
		}
	}()

	// Queue of requests.
	burstyRequests := make(chan int, 5)

	// Add 5 requests.
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}

	// No more requests will be sent.
	close(burstyRequests)

	// Process requests.
	for req := range burstyRequests {

		// Take a token from the bucket.
		//
		// If a token exists:
		//     proceed immediately.
		//
		// If bucket is empty:
		//     wait until refill goroutine
		//     adds another token.
		<-burstyLimiter

		fmt.Println("request", req, time.Now())
	}
}
