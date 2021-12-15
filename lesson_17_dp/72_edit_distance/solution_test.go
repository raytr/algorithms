package edit_distance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 3, minDistance("horse", "ros"))
}
