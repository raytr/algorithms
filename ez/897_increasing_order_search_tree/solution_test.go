package increasing_order_search_tree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var input1 = &TreeNode{
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
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	},
}

var output1 = &TreeNode{
	Val: 1,
	Right: &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
						Right: &TreeNode{
							Val: 7,
							Right: &TreeNode{
								Val: 8,
								Right: &TreeNode{
									Val: 9,
								},
							},
						},
					},
				},
			},
		},
	},
}

var input2 = &TreeNode{
	Val: 5,
	Left: &TreeNode{
		Val: 3,
	},
	Right: &TreeNode{
		Val: 6,
	},
}

var output2 = &TreeNode{
	Val: 3,
	Right: &TreeNode{
		Val: 5,
		Right: &TreeNode{
			Val: 6,
		},
	},
}

func TestSolution(t *testing.T) {
	tests := []struct {
		name   string
		input  *TreeNode
		output *TreeNode
	}{
		{"[5,3,6,2,4,null,8,1,null,null,null,7,9]", input1, output1},
		{"[5,3,6]", input2, output2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, increasingBST(tt.input), tt.output)
		})
	}
}
