package average_of_levels_in_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/require"
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
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 20,
		},
	}
)

func TestAverageOfLevels(t *testing.T) {
	testcases := []struct {
		description string
		root        *TreeNode
		want        []float64
	}{
		{"root = [3,9,20,null,null,15,7]", root1, []float64{3.00000, 14.50000, 11.00000}},
		{"root = [3,9,20,15,7]", root2, []float64{3.00000, 14.50000, 11.00000}},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			require.Equal(t, tc.want, averageOfLevels(tc.root))
		})
	}
}
