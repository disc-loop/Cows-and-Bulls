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

/*
func isStringPresentIn(slice []string, s string) bool {
	var result bool

	for _, v := range slice {
		if v == s {
			result = true
		}
	}

	return result
}

func compareSharedElem(first []string, second []string) map[string]string {
	shared := make(map[string]string)

	for _, fVal := range first {
		for _, sVal := range second {
			if fVal == sVal {
				shared[fVal] = sVal
			}
		}
	}

	return shared
}
*/

func Compare(secret string, guess string) Accuracy {
	var s []string = strings.Split(secret, "")
	var g []string = strings.Split(guess, "")
	exact := 0
	near := 0

	/* First case - Exact matches
	for gIdx, gVal := range g {
		if gVal == s[gIdx] {
			exact++
		}
	}
	*/

	// Second case - Near matches (now combined)
	// Does not work for all cases as it counts duplicate matches
	for gIdx, gVal := range g {
		for _, sVal := range s {
			// Count exacts
			if g[gIdx] == s[gIdx] {
				exact++
			}
			if gVal == sVal {
				fmt.Println(gIdx, gVal)
				near++
				break
			}
		}
	}

	// To match nears:
	// it cannot be a previously matched near (number and position)
	// it cannot be a previously matched exact (number and position)
	// it cannot be an exact (number and position)
	// only matched if these conditions are met
	// I need something that tracks the previously matched digits with their positions
	// Map or slice?

	// 	New rules:
	// - An exact match means that a digit in the guess is in the same place as that of the secret.
	// e.g. Secret = 1423, Guess = 1099, Exact matches = 1 (the first digit of the guess matches the first digit of the secret)
	// - A near match means that a digit in the guess is present within the secret, but only if that digit has not already been near matched or isn't an exact match

	/*
			for i, sVal := range s {
				// Counts exact matches
				if g[i] == sVal {
					exact++
					// Counts the matched digits
					m = append(m, sVal) // Keys must be unique, else they will be overwritten
				}
			}

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
