package guess

import (
	"fmt"

	"github.com/Tom/CowsAndBulls/pkg/input"
)

type Accuracy struct {
	Exact int
	Near  int
}

func (accuracy Accuracy) ToString() string {
	if accuracy == (Accuracy{}) {
		return "No matches!"
	}

	return fmt.Sprintf("Bulls: %v, Cows: %v", accuracy.Exact, accuracy.Near)
}

func Compare(secret string, guess string) (Accuracy, error) {
	// Input validation
	if !input.CorrectLength(secret, guess) {
		errString := fmt.Sprintf("incorrect string length for guess. Should be %v characters long.", len(secret))
		return Accuracy{}, fmt.Errorf(errString)
	}

	var s []byte = []byte(secret)
	var g []byte = []byte(guess)
	len := len(secret)
	accuracy := Accuracy{}

	// First case - Exact matches
	for i := 0; i < len; i++ {
		if s[i] == g[i] {
			accuracy.Exact++
			s[i] = 'x'
		}
	}

	// Second case - Near matches
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if g[i] == s[j] {
				accuracy.Near++
				s[j] = 'x'
				break
			}
		}
	}

	return accuracy, nil
}
