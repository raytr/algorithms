package longest_well_performing_interval

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 3, longestWPI([]int{9, 9, 6, 0, 6, 6, 9}))
	assert.Equal(t, 0, longestWPI([]int{6, 6, 6}))
	assert.Equal(t, 3, longestWPI([]int{6, 9, 9}))
	assert.Equal(t, 3, longestWPI([]int{6, 6, 6, 9, 9, 6, 6}))
	assert.Equal(t, 7, longestWPI([]int{6, 6, 9, 9, 9, 9, 6, 6}))
	assert.Equal(t, 3, longestWPI([]int{8, 10, 6, 16, 5}))
	assert.Equal(t, 3, longestWPI([]int{10, 5, 10, 6, 8, 6, 5, 7, 5, 8, 6}))
}
