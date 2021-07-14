package generate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecret(t *testing.T) {
	s := Secret(2)
	assert.Equal(t, "2716", s)
}
