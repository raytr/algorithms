package path_sum_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var input = &TreeNode{
	Val: 5,
	Left: &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 11,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
	},
	Right: &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 13,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	},
}

func TestSolution(t *testing.T) {
	assert.Equal(t, [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}, pathSum(input, 22))
}
