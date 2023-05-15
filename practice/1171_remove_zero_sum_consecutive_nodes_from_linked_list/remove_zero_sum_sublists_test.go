package remove_zero_sum_consecutive_nodes_from_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var head1 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: -3,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	},
}

var expectation1 = &ListNode{
	Val: 3,
	Next: &ListNode{
		Val: 1,
	},
}

var head2 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: -3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	},
}

var expectation2 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
		},
	},
}

func TestRemoveZeroSumSublists(t *testing.T) {
	tests := []struct {
		name        string
		head        *ListNode
		expectation *ListNode
	}{
		{"head = [1,2,-3,3,1]", head1, expectation1},
		{"head = [1,2,3,-3,4]", head2, expectation2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation, removeZeroSumSublists(tt.head))
		})
	}
}
