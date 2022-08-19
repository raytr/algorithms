package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//go over 2 lists and compute the sum, if sum > 9, remember is true
	// next node, if remember is true => +1 to sum
	res := new(ListNode)
	cur := res
	sum := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{
			Val: sum % 10,
		}
		cur = cur.Next
		sum /= 10
	}

	if sum > 0 {
		cur.Next = &ListNode{Val: sum}
	}
	return res.Next
}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	remember := 0
//	output := new(ListNode)
//	cur := output
//
//	for l1 == nil {
//		sum := l1.Val + l2.Val + remember
//		if sum > 9 {
//			remember = 1
//		} else {
//			remember = 0
//		}
//
//		cur = &ListNode{
//			Val: sum % 10,
//		}
//		l1 = l1.Next
//		l2 = l2.Next
//		cur = cur.Next
//	}
//	return output
//}
