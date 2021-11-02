package search_in_rotated_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, 4, search([]int{4,5,6,7,0,1,2},0))
	assert.Equal(t, 2, search([]int{5, 1, 3}, 3))
}
