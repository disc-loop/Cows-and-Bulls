package input

import "errors"

type ValidGuess string

func ValidateGuess(secret string, guess string) (ValidGuess, error) {
	if correctLength(secret, guess) {
		return ValidGuess(guess), nil
	}
	//Add correct character set
	return "", errors.New("invalid guess")
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

// // Validates guess and secret are of same character set
// func CorrectCharacterSet(guess string) bool {
// 	return true
// }
// // Could add another param for character set
