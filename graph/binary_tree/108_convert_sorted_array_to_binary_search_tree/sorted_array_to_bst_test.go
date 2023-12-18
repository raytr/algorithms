package convert_sorted_array_to_binary_search_tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	expectation1 = &TreeNode{
		Left: &TreeNode{
			Right: &TreeNode{
				Val: -3,
			},
			Val: -10,
		},
		Right: &TreeNode{
			Right: &TreeNode{
				Val: 9,
			},
			Val: 5,
		},
		Val: 0,
	}

	expectation2 = &TreeNode{
		Right: &TreeNode{
			Val: 3,
		},
		Val: 1,
	}
)

func TestSortedArrayToBST(t *testing.T) {
	testsCases := []struct {
		name string
		nums []int
		want *TreeNode
	}{
		{"nums = [-10,-3,0,5,9]", []int{-10, -3, 0, 5, 9}, expectation1},
		{"nums = [1,3]", []int{1, 3}, expectation2},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, sortedArrayToBST(tt.nums))
		})
	}
}
