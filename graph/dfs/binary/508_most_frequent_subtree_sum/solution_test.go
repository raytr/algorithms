package most_frequent_subtree_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
