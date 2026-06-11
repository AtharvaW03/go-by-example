package main

import "fmt"

// Variadic function.
// The ...int means this function can accept any number of ints.
// Inside the function, nums is actually a []int slice.
func sum(nums ...int) {
	// Print the slice of numbers received.
	fmt.Print(nums, " ")

	total := 0

	// Loop through every element in the nums slice.
	for _, num := range nums {
		total += num
	}

	// Print the final sum.
	fmt.Println(total)
}

func main() {
	// Pass 2 arguments.
	sum(1, 2)

	// Pass 3 arguments.
	sum(1, 2, 3)

	// Create a slice.
	nums := []int{1, 2, 3, 4}

	// Expand/unpack the slice into individual arguments.
	// Equivalent to: sum(1, 2, 3, 4)
	sum(nums...)
}