package main

import "fmt"

// Define a struct type named person.
// Similar to a simple Java class with fields.
type person struct {
	name string
	age  int
}

// Constructor-like function.
// Creates a person and returns a pointer to it.
func newPerson(name string) *person {
	// Create a person with only the name set.
	p := person{name: name}

	// Set age after creation.
	p.age = 42

	// Return the address of p (type: *person).
	return &p
}

func main() {
	// Create a person using positional fields.
	// Order must match the struct definition.
	fmt.Println(person{"Bob", 20})

	// Create a person using named fields.
	// More readable and safer.
	fmt.Println(person{name: "Alice", age: 21})

	// Unspecified fields get zero values.
	// age defaults to 0.
	fmt.Println(person{name: "Fred"})

	// Create a struct and immediately take its address.
	// Result type is *person.
	fmt.Println(&person{name: "Ann", age: 44})

	// Call constructor-like function.
	// Returns *person.
	fmt.Println(newPerson("Jon"))

	// Create a normal person value.
	s := person{name: "Sean", age: 55}

	// Access a field.
	fmt.Println(s.name)

	// Take the address of s.
	// sp is a *person.
	sp := &s

	// Go automatically dereferences pointers for struct fields.
	// Equivalent to (*sp).age.
	fmt.Println(sp.age)

	// Modify the original struct through the pointer.
	sp.age = 51

	// Prints updated value.
	fmt.Println(sp.age)

	// Anonymous struct:
	// define the struct type and create an instance at the same time.
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)
}
