package input

import (
	"fmt"
	"regexp"
)

type ValidGuess string

type lengthError struct {
	length int
}

func (error lengthError) Error() string {
	return fmt.Sprintf("incorrect length for guess. Guess should be of length %v", error.length)
}

type characterError struct{}

func (error characterError) Error() string {
	return "invalid characters. Guess should be numeric only"
}

func ValidateGuess(secret string, guess string) (ValidGuess, error) {
	if !correctLength(secret, guess) {
		return "", lengthError{length: len(secret)}
	}
	if !correctCharacterSet(guess) {
		return "", characterError{}
	}
	//Add correct character set
	return ValidGuess(guess), nil
}

// Validates guess and secret are same size
func correctLength(secret string, guess string) bool {
	var s []byte = []byte(secret)
	var g []byte = []byte(guess)

	if len(s) == len(g) {
		return true
	} else {
		return false
	}
}

// Validates guess and secret are of same character set
func correctCharacterSet(guess string) bool {
	set := regexp.MustCompile(`\A\d+\z`)
	return set.MatchString(guess)
}

// Could add another param for character set
