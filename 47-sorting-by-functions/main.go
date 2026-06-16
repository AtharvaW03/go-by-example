package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {

	// Slice of strings.
	fruits := []string{"peach", "banana", "kiwi"}

	// Comparison function.
	//
	// Compare strings based on their length.
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// Sort fruits by length.
	slices.SortFunc(fruits, lenCmp)

	fmt.Println(fruits)

	// Person type.
	type Person struct {
		name string
		age  int
	}

	// Slice of people.
	// Note: if the Person struct is large,
	// you may want the slice to contain *Person instead
	// and adjust the sorting function accordingly.
	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}

	// Sort people by age.
	slices.SortFunc(
		people,
		func(a, b Person) int {

			// Compare ages.
			return cmp.Compare(a.age, b.age)
		},
	)

	fmt.Println(people)
}
