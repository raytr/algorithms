package n_ary_tree_level_order_traversal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	root1 = &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}

	root2 = &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 2,
			},
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 6,
					},
					{
						Val: 7,
						Children: []*Node{
							{
								Val: 11,
								Children: []*Node{
									{
										Val: 14,
									},
								},
							},
						},
					},
				},
			},
			{
				Val: 4,
				Children: []*Node{
					{
						Val: 8,
						Children: []*Node{{
							Val: 12,
						},
						},
					},
				},
			},
			{
				Val: 5,
				Children: []*Node{
					{
						Val: 9,
						Children: []*Node{
							{
								Val: 13,
							},
						},
					},
					{
						Val: 10,
					},
				},
			},
		},
	}
)

func TestLevelOrder(t *testing.T) {

	testCases := []struct {
		name string
		root *Node
		want [][]int
	}{
		{"root = [1,null,3,2,4,null,5,6]", root1, [][]int{{1}, {3, 2, 4}, {5, 6}}},
		{"root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]", root2, [][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14}}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, levelOrder(tt.root))
		})
	}
}
