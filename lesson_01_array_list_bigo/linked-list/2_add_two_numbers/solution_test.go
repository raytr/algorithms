package add_two_numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var l11 = &ListNode{
	Val: 2,
	Next: &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 3,
		},
	},
}

var l12 = &ListNode{
	Val: 5,
	Next: &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 4,
		},
	},
}

var expectation1 = &ListNode{
	Val: 7,
	Next: &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 8,
		},
	},
}

var l21 = &ListNode{
	Val: 9,
	Next: &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
						},
					},
				},
			},
		},
	},
}

var l22 = &ListNode{
	Val: 9,
	Next: &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
			},
		},
	},
}

var expectation2 = &ListNode{
	Val: 8,
	Next: &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 1,
							},
						},
					},
				},
			},
		},
	},
}

func TestSolution(t *testing.T) {
	//assert.Equal(t, expectation1, addTwoNumbers(l11,l12))
	assert.Equal(t, expectation2, addTwoNumbers(l21, l22))
}
