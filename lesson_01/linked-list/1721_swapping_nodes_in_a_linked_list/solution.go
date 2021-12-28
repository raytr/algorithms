package swapping_nodes_in_a_linked_list

/*
	problem: https://leetcode.com/problems/swapping-nodes-in-a-linked-list/submissions/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	p1 := head
	for i := 1; i < k; i++ {
		if p1.Next == nil {
			return nil
		}
		p1 = p1.Next
	}

	temp := p1
	p2 := head

	for p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	temp.Val, p2.Val = p2.Val, temp.Val
	return head
}
