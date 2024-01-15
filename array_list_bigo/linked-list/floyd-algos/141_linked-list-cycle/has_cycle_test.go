package linked_list_cycle

import "testing"

//1 -> 2 -> 3

var input1 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  3,
			Next: nil,
		},
	},
}

func TestName(t *testing.T) {

}
