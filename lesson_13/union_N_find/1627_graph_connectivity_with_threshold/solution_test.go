package graph_connectivity_with_threshold

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, []bool{false, false, true}, areConnected(6, 2, [][]int{{1, 4}, {2, 5}, {3, 6}}))
	assert.Equal(t, []bool{true, true, true, true, true}, areConnected(6, 0, [][]int{{4, 5}, {3, 4}, {3, 2}, {2, 6}, {1, 3}}))
	assert.Equal(t, []bool{false, false, false, false, false}, areConnected(5, 1, [][]int{{4, 5}, {4, 5}, {3, 2}, {2, 3}, {3, 4}}))
}
