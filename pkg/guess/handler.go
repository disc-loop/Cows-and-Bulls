package guess

import (
	"fmt"
	"strings"
)

// One solution for returning two integers.
type Accuracy struct {
	Exact int
	Near  int
}

// Another would be to return two values and have the response creator
// handle the two values. Possibly a more golike approach.
// func Compare(guess [4]int, secret [4]int) (int, int) {}

type matchDetails struct {
	value string
	exact string
}

func Compare(secret string, guess string) Accuracy {
	var s []string = strings.Split(secret, "")
	var g []string = strings.Split(guess, "")
	var matched = make(map[int]matchDetails) // key = position, matchDetails
	exact := 0
	near := 0
	duplicates := false

	// Testing
	fmt.Println(s, g)

	for i, v := range s {
		// Counts exact matches
		if g[i] == v {
			exact++
			// Counts the matched digits
			matched[i] = matchDetails{v, "exact"} // Keys must be unique, else they will be overwritten
		}
	}
	for _, v := range matched {
		// Adds near matches if they have not been matched already
		for _, digit := range g {
			if digit == v.value {
				duplicates = true
			}
		}
		if !duplicates {
			near++
		}
	}

	

	// Testing
	fmt.Println(matched)

	// I need a function that takes the secret (string) and converts it
	// to an array of integers so that I can check for cows (if a number in guess matches)
	// and for bulls (if number in guess matches and is the same index).

	// Matched should be a bool. If false, a position can be matched. If true, it cannot be matched again.

	return Accuracy{Exact: exact, Near: 0}
}
