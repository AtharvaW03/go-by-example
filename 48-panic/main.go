package main

import (
	"os"
	"path/filepath"
)

func main() {

	// Trigger a panic immediately.
	//
	// The program stops here and prints
	// a stack trace.
	panic("a problem")

	// This code never executes because the
	// panic above terminates the program.
	path := filepath.Join(os.TempDir(), "file")

	// Try to create a file.
	_, err := os.Create(path)

	// If file creation fails, panic with
	// the returned error.
	if err != nil {
		panic(err)
	}
}
