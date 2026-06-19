package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Struct without JSON tags.
// JSON keys will be "Page" and "Fruits".
type response1 struct {
	Page   int
	Fruits []string
}

// Struct with JSON tags.
// JSON keys will be "page" and "fruits".
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// -----------------------------
	// Basic types -> JSON
	// -----------------------------

	// Convert bool to JSON.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	// Convert int to JSON.
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	// Convert float to JSON.
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	// Convert string to JSON.
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// -----------------------------
	// Slice -> JSON
	// -----------------------------

	slcD := []string{"apple", "peach", "pear"}

	// Convert slice to JSON array.
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	// -----------------------------
	// Map -> JSON
	// -----------------------------

	mapD := map[string]int{
		"apple":   5,
		"lettuce": 7,
	}

	// Convert map to JSON object.
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// -----------------------------
	// Struct -> JSON
	// -----------------------------

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}

	// Convert struct to JSON.
	// Uses field names Page and Fruits.
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}

	// Convert struct to JSON.
	// Uses tags: page and fruits.
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// -----------------------------
	// JSON -> map
	// -----------------------------

	// Raw JSON data.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// Generic JSON object.
	var dat map[string]interface{}

	// Parse JSON into map.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	fmt.Println(dat)

	// Type assertion:
	// Tell Go this value is a float64.
	num := dat["num"].(float64)
	fmt.Println(num)

	// JSON array becomes []interface{}.
	strs := dat["strs"].([]interface{})

	// First element is a string.
	str1 := strs[0].(string)
	fmt.Println(str1)

	// -----------------------------
	// JSON -> Struct
	// -----------------------------

	str := `{"page": 1, "fruits": ["apple", "peach"]}`

	// Empty struct.
	res := response2{}

	// Parse JSON into struct.
	_ = json.Unmarshal([]byte(str), &res)

	fmt.Println(res)

	// Access struct field normally.
	fmt.Println(res.Fruits[0])

	// -----------------------------
	// Encoder
	// -----------------------------

	// Encoder writes JSON directly
	// to the provided writer.
	enc := json.NewEncoder(os.Stdout)

	d := map[string]int{
		"apple":   5,
		"lettuce": 7,
	}

	// Write JSON to stdout.
	_ = enc.Encode(d)

	// -----------------------------
	// Decoder
	// -----------------------------

	// Decoder reads JSON from a stream.
	dec := json.NewDecoder(strings.NewReader(str))

	res1 := response2{}

	// Read JSON into struct.
	_ = dec.Decode(&res1)

	fmt.Println(res1)
}
