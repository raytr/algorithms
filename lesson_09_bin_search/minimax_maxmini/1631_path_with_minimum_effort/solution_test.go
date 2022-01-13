package path_with_minimum_effort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 2, minimumEffortPath([][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}))
	assert.Equal(t, 0, minimumEffortPath([][]int{{1}}))
	assert.Equal(t, 0, minimumEffortPath([][]int{{9}}))
	assert.Equal(t, 0, minimumEffortPath([][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}))
}
