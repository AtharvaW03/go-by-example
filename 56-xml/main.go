package main

import (
	"encoding/xml"
	"fmt"
)

// Plant represents an XML element named <plant>.
type Plant struct {

	// Root XML tag name.
	XMLName xml.Name `xml:"plant"`

	// XML attribute:
	// <plant id="27">
	Id int `xml:"id,attr"`

	// XML element:
	// <name>Coffee</name>
	Name string `xml:"name"`

	// Repeated XML elements:
	// <origin>Ethiopia</origin>
	// <origin>Brazil</origin>
	Origin []string `xml:"origin"`
}

// Custom string representation.
func (p Plant) String() string {
	return fmt.Sprintf(
		"Plant id=%v, name=%v, origin=%v",
		p.Id,
		p.Name,
		p.Origin,
	)
}

func main() {

	// Create plant.
	coffee := &Plant{
		Id:   27,
		Name: "Coffee",
	}

	// Multiple origins.
	coffee.Origin = []string{
		"Ethiopia",
		"Brazil",
	}

	// -----------------------------
	// Go struct -> XML
	// -----------------------------

	// Convert struct to formatted XML.
	out, _ := xml.MarshalIndent(
		coffee,
		" ",
		"  ",
	)

	fmt.Println(string(out))

	// Print XML declaration header
	// plus XML content.
	fmt.Println(xml.Header + string(out))

	// -----------------------------
	// XML -> Go struct
	// -----------------------------

	var p Plant

	// Parse XML into struct.
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}

	// Uses Plant.String().
	fmt.Println(p)

	// -----------------------------
	// Another plant
	// -----------------------------

	tomato := &Plant{
		Id:   81,
		Name: "Tomato",
	}

	tomato.Origin = []string{
		"Mexico",
		"California",
	}

	// -----------------------------
	// Nested XML example
	// -----------------------------

	type Nesting struct {

		// Root tag:
		// <nesting>
		XMLName xml.Name `xml:"nesting"`

		// Create nested structure:
		//
		// parent
		//   child
		//     plant
		Plants []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}

	// Add both plants.
	nesting.Plants = []*Plant{
		coffee,
		tomato,
	}

	// Convert nested struct to XML.
	out, _ = xml.MarshalIndent(
		nesting,
		" ",
		"  ",
	)

	fmt.Println(string(out))
}
