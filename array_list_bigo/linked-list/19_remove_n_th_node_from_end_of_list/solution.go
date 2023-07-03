package remove_n_th_node_from_end_of_list

/**
problem: https://leetcode.com/problems/remove-nth-node-from-end-of-list
time complexity: O(n)
*/

type listNode struct {
	val  int
	next *listNode
}

func removeNthFromEnd(A *listNode, B int) *listNode {
	//A = &listNode{
	//	val:  0,
	//	next: A,
	//}
	//clone := A
	//node, _ := getOrder(clone, B)
	//node.next = node.next.next
	//return A.next

	f := A
	count := 0
	for count < B {
		if f == nil {
			return A.next
		}
		f = f.next
		count++
	}

	//edge cases: B == length of the LL
	if f == nil {
		return A.next
	}

	s := A
	for f.next != nil {
		f = f.next
		s = s.next
	}

	s.next = s.next.next

	return A
}

func getOrder(head *listNode, n int) (*listNode, int) {
	if head == nil {
		return nil, 0
	}

	node, order := getOrder(head.next, n)
	if order == n+1 {
		return node, order
	}
	return head, order + 1
}
