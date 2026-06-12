package main

import "fmt"

// Base struct.
type base struct {
	num int
}

// Method belonging to base.
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// container embeds base.
// This is called struct embedding.
type container struct {
	base // embedded struct
	str  string
}

func main() {

	// Create a container.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// Because base is embedded, num is promoted.
	// This works:
	//   co.num
	// instead of:
	//   co.base.num
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Direct access through the embedded struct.
	fmt.Println("also num:", co.base.num)

	// describe() belongs to base, not container.
	// But embedding promotes methods too, so this works:
	//   co.describe()
	// Equivalent to:
	//   co.base.describe()
	fmt.Println("describe:", co.describe())

	// Interface requiring a describe() method.
	type describer interface {
		describe() string
	}

	// container satisfies describer because
	// it gets describe() from the embedded base.
	var d describer = co

	// Call describe() through the interface.
	fmt.Println("describer:", d.describe())
}
