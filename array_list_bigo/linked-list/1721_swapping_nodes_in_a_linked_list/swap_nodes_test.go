package swapping_nodes_in_a_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var input = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
	},
}

var want = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
	},
}

func TestSolution(t *testing.T) {
	assert.Equal(t, expectation, swapNodes(input, 2))
}
