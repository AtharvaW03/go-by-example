package main

import "fmt"

// Recursive factorial function.
func fact(n int) int {
	// Base case.
	if n == 0 {
		return 1
	}

	// Recursive case.
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7)) // 5040

	// Declare a variable that can hold a function.
	var fib func(n int) int

	// Assign a recursive anonymous function.
	fib = func(n int) int {
		// Base cases.
		if n < 2 {
			return n
		}

		// Recursive case.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7)) // 13
}
