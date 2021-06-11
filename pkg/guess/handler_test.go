package guess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareReturnsAllExact(t *testing.T) {
	result := Compare("1111", "1111")

	assert.Equal(t, Accuracy{Exact: 4}, result)
}

func TestCompareReturnsNearsCorrectly(t *testing.T) {
	result := Compare("1234", "4321")
	result2 := Compare("1134", "0001")

	assert.Equal(t, Accuracy{Near: 4}, result)
	assert.Equal(t, Accuracy{Near: 1}, result2)
}

func TestCompareReturnsExpectedAccuracy(t *testing.T) {
	result1 := Compare("1234", "1822")
	result2 := Compare("1224", "1812")

	assert.Equal(t, Accuracy{Exact: 1, Near: 1}, result1)
	assert.Equal(t, Accuracy{Exact: 1, Near: 1}, result2)
}

func TestCompareReturnsEmptyOnNoMatches(t *testing.T) {
	result := Compare("1111", "9999")

	assert.Equal(t, Accuracy{}, result)
}
func TestCompareDoesNotDuplicateMatches(t *testing.T) {
	result1 := Compare("1134", "1011")
	result2 := Compare("1134", "0011")

	assert.Equal(t, Accuracy{Exact: 1, Near: 1}, result1)
	assert.Equal(t, Accuracy{Exact: 0, Near: 2}, result2)
}

func TestCompareHandlesLongerNumbers(t *testing.T) {
	result := Compare("123456789", "987654321")

	assert.Equal(t, Accuracy{Exact: 1, Near: 8}, result)
}
