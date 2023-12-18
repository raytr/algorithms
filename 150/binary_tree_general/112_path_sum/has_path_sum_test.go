package path_sum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	input1 = &TreeNode{
		Left: &TreeNode{
			Left: &TreeNode{
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
				Val: 11,
			},
			Val: 4,
		},
		Right: &TreeNode{
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Right: &TreeNode{
					Val: 1,
				},
				Val: 4,
			},
			Val: 8,
		},
		Val: 5,
	}

	input2 = &TreeNode{
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
		Val: 1,
	}
)

func TestHasPathSum(t *testing.T) {
	testCases := []struct {
		description string
		root        *TreeNode
		targetSum   int
		want        bool
	}{
		{"root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22", input1, 22, true},
		{"root = [1,2,3], targetSum = 5", input2, 5, false},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			require.Equal(t, tt.want, hasPathSum(tt.root, tt.targetSum))
		})
	}
}
