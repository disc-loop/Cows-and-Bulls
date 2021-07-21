package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateGuessWorks(t *testing.T) {
	s1 := "1234"
	s2 := "4321"

	g, err := ValidateGuess(s1, s2)

	assert.NoError(t, err)
	assert.Equal(t, ValidGuess("4321"), g)
}

func TestValidateGuessError(t *testing.T) {
	secret := "1234"
	tests := []struct {
		input    string
		expected error
		name     string
	}{{
		input:    "12345",
		expected: lengthError{len(secret)},
		name:     "OnIncorrectLength",
	}, {
		input:    "1a23",
		expected: characterError{},
		name:     "OnIncorrectCharacterSet",
	}}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			_, err := ValidateGuess(secret, v.input)

			assert.Equal(t, v.expected, err)
		})
	}
}
