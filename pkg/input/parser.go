package input

// Validates that guess and secret are same size

func CorrectLength(secret string, guess string) bool {
	var s []byte = []byte(secret)
	var g []byte = []byte(guess)

	if len(s) == len(g) {
		return true
	} else {
		return false
	}
}
