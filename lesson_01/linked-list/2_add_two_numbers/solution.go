package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//reverse 2 lists
	l1 = reverseList(l1)
	l2 = reverseList(l2)

	//go over 2 lists and compute the sum, if sum > 9, remember is true
	// next node, if remember is true => +1 to sum
	res := &ListNode{}
	cur := res
	remember := 0
	for l1 != nil || l2 != nil {
		cur.Next = &ListNode{}
		cur = cur.Next
		sum := 0
		if l1 == nil {
			sum = l2.Val + remember
			l2 = l2.Next
		} else if l2 == nil {
			sum = l1.Val + remember
			l1 = l1.Next
		} else {
			sum = l1.Val + l2.Val + remember
			l1 = l1.Next
			l2 = l2.Next
		}

		//handle when sum > 9
		if remember > 0 {
			remember--
		}
		if sum > 9 {
			sum -= 10
			remember++
		}
		cur.Val = sum
	}
	if remember > 0 {
		cur.Next = &ListNode{
			Val: 1,
		}

	}
	return res.Next
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
