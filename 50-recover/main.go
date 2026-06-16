package main

import "fmt"

// Function that always panics.
func mayPanic() {
	panic("a problem")
}

func main() {

	// Deferred function runs if a panic occurs.
	defer func() {

		// recover() returns the panic value.
		// If no panic is happening, it returns nil.
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	// Triggers a panic.
	mayPanic()

	// This line never executes because
	// the panic interrupts normal flow.
	fmt.Println("After mayPanic()")
}
