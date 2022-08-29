package four_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, [][]int{{-2,-1,1,2},{-2,0,0,2},{-1,0,0,1}}, fourSum([]int{1,0,-1,0,-2,2},0))
	//assert.Equal(t, [][]int{{2,2,2,2}}, fourSum([]int{2,2,2,2,2},8))
	//assert.Equal(t, [][]int{{2,2,2,2}}, fourSum([]int{2,2,2,2,2},8))
	assert.Equal(t, [][]int{{-2, -1, 1, 2}, {-1, -1, 1, 1}}, fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0))
}
