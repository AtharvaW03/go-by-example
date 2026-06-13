package main

import (
	"errors"
	"fmt"
)

// Function returning:
// 1. a result
// 2. an error
func f(arg int) (int, error) {

	// Simulate a failure case.
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	// Success: return result and nil error.
	return arg + 3, nil
}

// Sentinel (named) errors.
// Useful because callers can compare against them.
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

// Function that returns only an error.
func makeTea(arg int) error {

	// Out of tea.
	if arg == 2 {
		return ErrOutOfTea

		// Power failure.
	} else if arg == 4 {

		// Wrap ErrPower while preserving it.
		// errors.Is() can still find ErrPower later.
		return fmt.Errorf("making tea: %w", ErrPower)
	}

	// Success.
	return nil
}

func main() {

	// Test f() with two values.
	for _, i := range []int{7, 42} {

		// Initialization part:
		//     r, e := f(i)
		//
		// Condition part:
		//     e != nil
		//
		// Equivalent to:
		//
		// r, e := f(i)
		// if e != nil { ... }
		if r, e := f(i); e != nil {

			// Error path.
			fmt.Println("f failed:", e)

		} else {

			// Success path.
			fmt.Println("f worked:", r)
		}
	}

	// Go 1.22:
	// Generates 0,1,2,3,4
	for i := range 5 {

		// Create err and immediately check it.
		if err := makeTea(i); err != nil {

			// Check if the error is ErrOutOfTea.
			if errors.Is(err, ErrOutOfTea) {

				fmt.Println("We should buy new tea!")

				// Check if the error is ErrPower.
			} else if errors.Is(err, ErrPower) {

				fmt.Println("Now it is dark.")

			} else {

				// Some unexpected error.
				fmt.Printf("Unknown error: %s\n", err)
			}

			// Skip the success message.
			continue
		}

		// Only reached when makeTea returned nil.
		fmt.Println("Tea is ready!")
	}
}
