package most_frequent_subtree_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	input := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: -3,
		},
	}

	assert.Equal(t, []int{2, -3, 4}, findFrequentTreeSum(input))
}
