package palindrome_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//var input1 = &ListNode{
//	Val: 1,
//	Next: &ListNode{
//		Val: 2,
//		Next: &ListNode{
//			Val: 3,
//			Next: &ListNode{
//				Val: 2,
//				Next: &ListNode{
//					Val: 1,
//				},
//			},
//		},
//	},
//}

var input2 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,

		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
			},
		},
	},
}

func TestSolution(t *testing.T) {
	//assert.Equal(t, true, isPalindrome(input1))
	assert.Equal(t, true, isPalindrome(input2))
}
