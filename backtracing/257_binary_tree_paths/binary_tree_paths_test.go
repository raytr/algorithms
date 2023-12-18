package binary_tree_paths

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	root1 = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	root2 = &TreeNode{Val: 1}

	root3 = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 4},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
)

func TestBinaryTreePaths(t *testing.T) {
	testcases := []struct {
		description string
		root        *TreeNode
		want        []string
	}{
		{"root = [1,2,3,null,5]", root1, []string{"1->2->5", "1->3"}},
		{"root = [1]", root2, []string{"1"}},
		{"root = [1,2,3,4,5]", root3, []string{"1->2->4", "1->2->5", "1->3"}},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			require.ElementsMatch(t, testcase.want, binaryTreePaths(testcase.root))
		})
	}
}
