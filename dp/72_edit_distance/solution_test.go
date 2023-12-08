package edit_distance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, 3, minDistance("horse", "ros"))
	assert.Equal(t, 3, minDistanceRecursion("horse", "ros"))
}
