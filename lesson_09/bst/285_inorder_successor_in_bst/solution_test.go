package inorder_successor_in_bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	//root := &TreeNode{
	//	Val: 2,
	//	Left: &TreeNode{
	//		Val: 1,
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//	},
	//}
	//p := &TreeNode{
	//	Val: 1,
	//}

	root2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 6,
		},
	}
	//p2 := &TreeNode{
	//	Val: 2,
	//	Left: &TreeNode{
	//		Val: 1,
	//	},
	//}
	p2 := &TreeNode{
		Val: 6,
	}
	//res2 := nil
	//res2 := &TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{Val: 1,},
	//	},
	//	Right: &TreeNode{
	//		Val: 4,
	//	},
	//}
	//[5,3,6,2,4,null,null,1]
	//	6
	//	assert.Equal(t, root, inorderSuccessor(root, p))
	assert.Equal(t, nil, inorderSuccessor(root2, p2))
}
