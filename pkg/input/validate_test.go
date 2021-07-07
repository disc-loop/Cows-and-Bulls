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
	s1 := "1234"
	s2 := "43210"

	g, err := ValidateGuess(s1, s2)

	assert.Error(t, err)
	assert.Equal(t, ValidGuess(""), g)
}
