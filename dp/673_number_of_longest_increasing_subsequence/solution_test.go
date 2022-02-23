package number_of_longest_increasing_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, 1, findNumberOfLIS([]int{1}))
	//assert.Equal(t, 2, findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	//assert.Equal(t, 2, findNumberOfLIS([]int{1, 3, 5, 4}))
	//assert.Equal(t, 7, findNumberOfLIS([]int{2, 2, 2, 2, 2, 2, 2}))
	assert.Equal(t, 3, findNumberOfLIS([]int{1, 2, 4, 3, 5, 4, 7, 2}))
}
