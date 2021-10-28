package n_ary_tree_level_order_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	input := &Node{
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
	expectation := [][]int{{1}, {3, 2, 4}, {5, 6}}
	assert.Equal(t, expectation, levelOrder(input))
}
