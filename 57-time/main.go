package main

import (
	"fmt"
	"time"
)

func main() {

	// Shortcut for fmt.Println.
	p := fmt.Println

	// Current date and time.
	now := time.Now()
	p(now)

	// Create a specific date and time.
	//
	// Year:       2009
	// Month:      November
	// Day:        17
	// Hour:       20
	// Minute:     34
	// Second:     58
	// Nanosecond: 651387237
	// Timezone:   UTC
	then := time.Date(
		2009,
		11,
		17,
		20,
		34,
		58,
		651387237,
		time.UTC,
	)

	p(then)

	// Extract individual parts.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// Day of the week.
	p(then.Weekday())

	// Compare dates.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Time difference.
	//
	// Returns a Duration.
	diff := now.Sub(then)

	p(diff)

	// Convert duration into different units.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Add duration to a time.
	//
	// Should be approximately equal to now.
	p(then.Add(diff))

	// Subtract duration from a time.
	p(then.Add(-diff))
}
