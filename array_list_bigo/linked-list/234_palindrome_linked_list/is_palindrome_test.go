package palindrome_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//var input1 = &ListNode{
//	val: 1,
//	next: &ListNode{
//		val: 2,
//		next: &ListNode{
//			val: 3,
//			next: &ListNode{
//				val: 2,
//				next: &ListNode{
//					val: 1,
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
