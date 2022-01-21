package network

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, []int32{1, 3, 2, 3, 0, 0, 0, 0, 0}, distanceCount([]int32{9, 1, 4, 9, 0, 4, 8, 9, 0, 1}))
}
