package main

import (
	"errors"
	"fmt"
)

// Custom error type.
// Stores additional information about the error.
type argError struct {
	arg     int
	message string
}

// Implement the built-in error interface.
//
// Receiver:
//
//	e = receiver variable
//	*argError = receiver type (pointer to argError)
//
// Any type with Error() string satisfies the error interface.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

// Returns a result and an error.
func f(arg int) (int, error) {

	// Simulate a failure.
	if arg == 42 {

		// Create an argError struct and return a pointer to it.
		//
		// argError{...}  -> struct value
		// &argError{...} -> pointer to that struct (*argError)
		return -1, &argError{
			arg:     arg,
			message: "can't work with it",
		}
	}

	// Success.
	return arg + 3, nil
}

func main() {

	// Ignore the result, keep the error.
	_, err := f(42)

	// Try to extract an *argError from err.
	//
	// Read as:
	// "Is the error inside err actually an *argError?"
	//
	// If yes:
	//   ae = *argError
	//   ok = true
	//
	// If no:
	//   ok = false
	if ae, ok := errors.AsType[*argError](err); ok {

		// Access fields of the custom error.
		fmt.Println(ae.arg)
		fmt.Println(ae.message)

	} else {

		// Error was some other type.
		fmt.Println("err doesn't match argError")
	}
}
