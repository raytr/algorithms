package merge_intervals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, [][]int{{1, 6}, {8, 10}, {15, 18}}, merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	assert.Equal(t, [][]int{{0, 0}, {1, 4}}, merge([][]int{{1, 4}, {0, 0}}))
	assert.Equal(t, [][]int{{1, 10}}, merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
}
