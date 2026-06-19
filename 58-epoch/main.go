package main

import (
	"fmt"
	"time"
)

func main() {

	// Current date and time.
	now := time.Now()
	fmt.Println(now)

	// -----------------------------
	// Time -> Unix timestamp
	// -----------------------------

	// Seconds since Jan 1, 1970 UTC.
	fmt.Println(now.Unix())

	// Milliseconds since Jan 1, 1970 UTC.
	fmt.Println(now.UnixMilli())

	// Nanoseconds since Jan 1, 1970 UTC.
	fmt.Println(now.UnixNano())

	// -----------------------------
	// Unix timestamp -> Time
	// -----------------------------

	// Create time from Unix seconds.
	//
	// First argument = seconds
	// Second argument = nanoseconds
	fmt.Println(
		time.Unix(now.Unix(), 0),
	)

	// Create time from Unix nanoseconds.
	//
	// First argument = seconds
	// Second argument = nanoseconds
	fmt.Println(
		time.Unix(0, now.UnixNano()),
	)
}
