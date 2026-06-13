package main

import (
	"fmt"
	"time"
)

// Regular function.
func f(from string) {

	// Go 1.22:
	// Generates 0, 1, 2
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Normal function call.
	// main waits until f() finishes.
	f("direct")

	// Start f() in a new goroutine.
	// main does NOT wait.
	go f("goroutine")

	// Anonymous function.
	//
	// Parameter:
	//   msg string
	//
	// Argument passed:
	//   "going"
	//
	// Equivalent to:
	//
	// func printMsg(msg string) {
	//     fmt.Println(msg)
	// }
	//
	// go printMsg("going")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Give goroutines time to run.
	//
	// Without this, main may exit before
	// the goroutines get a chance to execute.
	time.Sleep(time.Second)

	fmt.Println("done")
}
