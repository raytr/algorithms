package diameter_of_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	root1 = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	root2 = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	root3 = &TreeNode{
		Val: 1,
	}

	root4 = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 6},
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
)

func TestSolution(t *testing.T) {

	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"[1,2,3,4,5]", root1, 3},
		{"[1,2]", root2, 1},
		{"[1]", root3, 0},
		{"[1,2,3,4,5, null, null,null,null,6]", root4, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, diameterOfBinaryTree(tt.root))
		})
	}
}
