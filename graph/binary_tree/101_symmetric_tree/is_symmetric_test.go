package symmetric_tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	input1 = &TreeNode{
		Left: &TreeNode{
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
			Val:   2,
		},
		Right: &TreeNode{
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
			Val:   2,
		},
		Val: 1,
	}
)

func TestIsSymmetric(t *testing.T) {
	testCases := []struct {
		name        string
		root        *TreeNode
		expectation bool
	}{
		{"root = [1,2,2,3,4,4,3]", input1, true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectation, isSymmetric(tt.root))
		})
	}
}
