package reverse_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {

	testcase1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	expection1 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	}

	testcase2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	expection2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}

	testCases := []struct {
		head      *ListNode
		expection *ListNode
	}{
		{testcase1, expection1},
		{testcase2, expection2},
	}

	for _, tc := range testCases {
		assert.Equal(t, reverseListByIterative(tc.head), tc.expection)
		assert.Equal(t, reverseListByRecursive(tc.head), tc.expection)
	}
}
