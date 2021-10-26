package number_of_islands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 1, numIslands([][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}))
	assert.Equal(t, 3, numIslands([][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}))
	assert.Equal(t, 2, numIslands([][]byte{
		{1, 0},
		{0, 1},
	}))
	assert.Equal(t, 1, numIslands([][]byte{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
}
