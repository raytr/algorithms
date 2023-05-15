package binary_tree_level_order_traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	root1 = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	root2 = &TreeNode{Val: 1}
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name        string
		root        *TreeNode
		expectation [][]int
	}{
		{"root = [3,9,20,null,null,15,7]", root1, [][]int{{3}, {9, 20}, {15, 7}}},
		{"root = [1]", root2, [][]int{{1}}},
		{"root = []", nil, [][]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation, levelOrder1(tt.root))
		})
	}
}
