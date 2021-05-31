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

// type matchDetails struct {
// 	value string
// 	matched bool
// }

func Compare(secret string, guess string) Accuracy {
	var s []string = strings.Split(secret, "")
	var g []string = strings.Split(guess, "")
	//var matched = make(map[int]matchDetails) // key = position, matchDetails
	var m []string
	exact := 0
	near := 0
	// duplicates := false

	// Testing
	fmt.Println(s, g)

	/*
		for i, sVal := range s {
			// Counts exact matches
			if g[i] == sVal {
				exact++
				// Counts the matched digits
				m = append(m, sVal) // Keys must be unique, else they will be overwritten
			}
		}
	*/

	// Checks for nears
	// Will only count digits if:
	// 1. The digit is present in the secret
	// 2. The digit has not been matched yet
	// i.e. Not an exact match (guess digit != secret digit)
	//  AND Not a previously counted near match (guess digit not in matched)

	dupes := 0
	for _, gVal := range g {
		// Check if gVal is in s
		for _, sVal := range s {
			// Check if gVal != sVal
			if gVal != sVal {
				// Check if gVal is not in m
				for _, mVal := range m {
					if gVal == mVal {
						dupes++
					}
				}
				if dupes == 0 {
					near++
					dupes = 0
				}
			}

		}

	}
	// Old near matcher
	/*
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
	*/

	// I need a function that takes the secret (string) and converts it
	// to an array of integers so that I can check for cows (if a number in guess matches)
	// and for bulls (if number in guess matches and is the same index).

	// Matched should be a bool. If false, a position can be matched. If true, it cannot be matched again.

	return Accuracy{Exact: exact, Near: near}
}
