package binary_tree_right_side_view

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	input := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 4},
		},
	}
	assert.Equal(t, []int{1, 3, 4}, rightSideView(input))
	assert.Equal(t, []int{}, rightSideView(nil))
	assert.Equal(t, []int{1}, rightSideView(&TreeNode{Val: 1}))
}
