package binary_tree_inorder_traversal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	input1 = &TreeNode{
		Left: nil,
		Right: &TreeNode{
			Left: &TreeNode{
				Val: 3,
			},
			Right: nil,
			Val:   2,
		},
		Val: 1,
	}

	input3 = &TreeNode{Val: 1}
)

func TestInorderTraversal(t *testing.T) {
	testCases := []struct {
		name string
		root *TreeNode
		want []int
	}{
		{"root = [1,null,2,3]", input1, []int{1, 3, 2}},
		{"root = []", nil, []int{}},
		{"root = [1]", input3, []int{1}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, inorderTraversal(tt.root))
		})
	}
}
