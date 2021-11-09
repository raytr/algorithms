package three_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
	assert.Equal(t, [][]int{{0, 0, 0}}, threeSum([]int{0, 0, 0}))
	assert.Equal(t, [][]int{{-2, 0, 2}, {-2, 1, 1}}, threeSum([]int{-2, 0, 1, 1, 2}))
	assert.Equal(t, [][]int{}, threeSum([]int{}))
	assert.Equal(t, [][]int{}, threeSum([]int{0}))
	assert.Equal(t, [][]int{{-5, 1, 4}, {-4, 0, 4}, {-4, 1, 3}, {-2, -2, 4}, {-2, 1, 1}, {0, 0, 0}}, threeSum([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}))
}
