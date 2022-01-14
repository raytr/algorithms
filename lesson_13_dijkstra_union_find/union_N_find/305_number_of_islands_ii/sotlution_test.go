package number_of_islands_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, []int{1, 1, 2, 3}, numIslands2(3, 3, [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}}))
	assert.Equal(t, []int{1, 1, 2, 2}, numIslands2(3, 3, [][]int{{0, 0}, {0, 1}, {1, 2}, {1, 2}}))
	assert.Equal(t, []int{1, 2, 2, 3, 4}, numIslands2(8, 4, [][]int{{0, 0}, {7, 1}, {6, 1}, {3, 3}, {4, 1}}))
	assert.Equal(t, []int{1, 1, 2, 3, 3, 3, 2, 2, 1, 1}, numIslands2(3, 3, [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}, {1, 0}, {0, 0}, {2, 2}, {1, 2}, {1, 1}, {0, 1}}))
}
