package find_minimum_in_rotated_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, 1, findMin([]int{3,4,5,1,2}))
	//assert.Equal(t, 1, findMin([]int{3,1,2}))
	assert.Equal(t, 1, findMin([]int{4, 5, 1, 2, 3}))
}
