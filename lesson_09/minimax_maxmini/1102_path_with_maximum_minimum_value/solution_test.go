package path_with_maximum_minimum_value

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, 4, maximumMinimumPath([][]int{{5, 4, 5},{1, 2, 6},{7, 4, 6}}))
	assert.Equal(t, 2, maximumMinimumPath([][]int{{2, 2, 1, 2, 2, 2}, {1, 2, 2, 2, 1, 2}}))
	//assert.Equal(t, 3, maximumMinimumPath([][]int{{3,4,6,3,4},{0,2,1,1,7},{8,8,3,2,7},{3,2,4,9,8},{4,1,2,0,0},{4,6,5,4,3}}))
}
