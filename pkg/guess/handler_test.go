package guess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Establish rules before filling out these test cases

func TestCompareReturnsAllExact(t *testing.T) {
	// Returns bulls correctly e.g. secret=1111 guess=1111 all bulls
	result := Compare("1111", "1111")

	assert.Equal(t, Accuracy{Exact: 4}, result)
}

func TestCompareReturnsNearsCorrectly(t *testing.T) {
	// Returns bulls correctly e.g. secret=1111 guess=1111 all bulls
	result := Compare("1234", "4321")

	assert.Equal(t, Accuracy{Near: 4}, result)
}

func TestCompareReturnsExpectedAccuracy(t *testing.T) {
	// Returns bulls correctly e.g. secret=1111 guess=1111 all bulls
	result := Compare("1234", "1823")

	assert.Equal(t, Accuracy{Exact: 1, Near: 2}, result)
}

func TestCompareReturnsEmptyOnNoMatches(t *testing.T) {
	// Returns bulls correctly e.g. secret=1111 guess=1111 all bulls
	result := Compare("1111", "9999")

	assert.Equal(t, Accuracy{}, result)
}
func TestCompareDoesNotDuplicateMatches(t *testing.T) {
	// Returns bulls correctly e.g. secret=1111 guess=1111 all bulls
	result1 := Compare("1134", "1011")
	result2 := Compare("1134", "0011")

	assert.Equal(t, Accuracy{Exact: 1, Near: 2}, result1)
	assert.Equal(t, Accuracy{Exact: 1, Near: 2}, result2)
}
