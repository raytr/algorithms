package shortest_path_in_binary_matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 2, shortestPathBinaryMatrix([][]int{{0, 1}, {1, 0}}))
	assert.Equal(t, -1, shortestPathBinaryMatrix([][]int{{1, 0}, {0, 1}}))
	assert.Equal(t, 4, shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}))
	assert.Equal(t, 25, shortestPathBinaryMatrix([][]int{{0, 1, 1, 1, 1, 1, 1, 1}, {0, 1, 1, 0, 0, 0, 0, 0}, {0, 1, 0, 1, 1, 1, 1, 0}, {0, 1, 0, 1, 1, 1, 1, 0}, {0, 1, 1, 0, 0, 1, 1, 0}, {0, 1, 1, 1, 1, 0, 1, 0}, {0, 0, 0, 0, 0, 1, 1, 0}, {1, 1, 1, 1, 1, 1, 1, 0}}))
}
