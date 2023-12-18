package merge_k_sorted_lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var ll1 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
		},
	},
}
var ll2 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
		},
	},
}
var ll3 = &ListNode{
	Val: 2,
	Next: &ListNode{
		Val: 6,
	},
}

var want = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
							},
						},
					},
				},
			},
		},
	},
}

func TestSolution(t *testing.T) {
	lists := []*ListNode{ll1, ll2, ll3}
	assert.Equal(t, expectation, mergeKLists(lists))
	assert.Equal(t, nil, mergeKLists(nil))
}
