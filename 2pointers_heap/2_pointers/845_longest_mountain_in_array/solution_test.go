package longest_mountain_in_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, 5, longestMountain([]int{2, 1, 4, 7, 3, 2, 5}))
	assert.Equal(t, 5, longestMountain([]int{1, 2, 3, 2, 1}))
}
