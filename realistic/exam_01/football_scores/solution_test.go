package football_scores

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, []int32{2, 3}, counts([]int32{1, 2, 3}, []int32{2, 4}))
	//assert.Equal(t, []int32{2, 4}, counts([]int32{1, 4, 2, 4}, []int32{3, 5}))
	assert.Equal(t, []int32{1, 0, 3, 4}, counts([]int32{2, 10, 5, 4, 8}, []int32{3, 1, 7, 8}))
}
