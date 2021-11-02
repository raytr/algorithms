package search_in_rotated_sorted_array_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, true, search([]int{2,5,6,0,0,1,2},0))
	//assert.Equal(t, false, search([]int{2,5,6,0,0,1,2},3))
	assert.Equal(t, true, search([]int{2, 2, 2, 2, 2, 1, 2}, 1))
	assert.Equal(t, true, search([]int{1, 2, 1}, 1))
	assert.Equal(t, true, search([]int{2, 1, 2}, 2))
}
