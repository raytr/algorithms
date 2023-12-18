package maximum_depth_of_binary_tree

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

	root2 = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}

	root3 = &TreeNode{
		Val: 1,
	}
)

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		name string
		root *TreeNode
		want int
	}{
		{"[3,9,20,null,null,15,7]", root1, 3},
		{"[1,null,2]", root2, 2},
		{"1", root3, 1},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxDepth(tt.root))
		})
	}
}
