package diameter_of_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 1, diameterOfBinaryTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}))
}
