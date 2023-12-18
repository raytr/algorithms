package merge_sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	testCases := []struct {
		Input *ListNode
		want  *ListNode
	}{
		//first test case
		{
			Input: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
			Expectation: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},

		//second test case
		{
			Input: &ListNode{
				Val: -1,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 0,
							},
						},
					},
				},
			},
			Expectation: &ListNode{
				Val: -1,
				Next: &ListNode{
					Val: 0,
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
			},
		},

		//nil test case
		{
			Input:       nil,
			Expectation: nil,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.want, sortList(tc.Input))
	}
}
