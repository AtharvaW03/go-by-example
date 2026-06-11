package main

import "fmt"

// intSeq returns a function.
// The returned function "remembers" the variable i.
func intSeq() func() int {
	i := 0 // This variable belongs to this specific closure.

	// Return an anonymous function.
	return func() int {
		i++      // Modify the captured variable.
		return i // Return the updated value.
	}
}

func main() {

	// Create a closure with its own i starting at 0.
	nextInt := intSeq()

	// Each call updates the same captured i.
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	// Create a completely new closure.
	// This gets its own separate i starting at 0.
	newInts := intSeq()

	// Uses its own i, not the one from nextInt.
	fmt.Println(newInts()) // 1
}