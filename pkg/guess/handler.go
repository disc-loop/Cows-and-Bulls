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

	for i, v := range s {
		// Counts exact matches
		if g[i] == v {
			exact++
			// Counts the matched digits
			m[i] = v // Keys must be unique, else they will be overwritten
		}
	}
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

	// Near++ if gVal is in s AND if gVal != matched (if gVal != in m)
	//(if g[i] != s[i] AND s[!i] has not been near matched) 
	for gPos, gVal := range g {
		
		// Check if gVal is in s
		for sPos, sVal := range s {
			
			// Check if gVal is in m AND if gVal != sVal
			for mPos, mVal := range m {
				if gVal == mVal {
					if gVal != sVal {
						near++
					}
				}
			}
		}
		
	} 
	// 

	// Testing
	fmt.Println(matched)

	// I need a function that takes the secret (string) and converts it
	// to an array of integers so that I can check for cows (if a number in guess matches)
	// and for bulls (if number in guess matches and is the same index).

	// Matched should be a bool. If false, a position can be matched. If true, it cannot be matched again.

	return Accuracy{Exact: exact, Near: 0}
}
