package response

import (
	"testing"

	"github.com/Tom/CowsAndBulls/pkg/guess"
	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	expected := "Bulls: 2, Cows: 2"
	accuracy := guess.Accuracy{Exact: 2, Near: 2}
	result := ToString(accuracy)

	assert.Equal(t, expected, result)
}

func TestToStringNoMatches(t *testing.T) {
	expected := "No matches!"
	accuracy := guess.Accuracy{Exact: 0, Near: 0}
	result := ToString(accuracy)

	assert.Equal(t, expected, result)
}