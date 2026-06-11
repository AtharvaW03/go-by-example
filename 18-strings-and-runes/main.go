package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "hello"

	// Number of bytes in the string.
	fmt.Println("Len:", len(s))

	// Print each byte in hexadecimal.
	for i := range len(s) {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// Number of runes (Unicode code points).
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// range automatically decodes runes.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")

	// Manually decode each rune from the UTF-8 string.
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])

		fmt.Printf("%#U starts at %d\n", runeValue, i)

		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'h' {
		fmt.Println("found h")
	}
}
