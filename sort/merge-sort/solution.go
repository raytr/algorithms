package merge_sort

//problem: https://leetcode.com/problems/sort-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	// when the linked list just has 1 node
	if head == nil || head.Next == nil {
		return head
	}

	//merge sort
	//at first, we find the middle node by floyd algorithm
	midNode := getTheMiddleNode(head)
	if midNode == nil {
		return head
	}

	right := midNode
	left := head
	for {
		if left.Next == midNode {
			left.Next = nil
			break
		}
		left = left.Next
	}
	left = head

	right = sortList(right)
	left = sortList(head)
	result := mergeTwoLists(left, right)
	return result
}

func getTheMiddleNode(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	result := new(ListNode)
	if l1.Val <= l2.Val {
		result.Val = l1.Val
		result.Next = mergeTwoLists(l1.Next, l2)
	} else {
		result.Val = l2.Val
		result.Next = mergeTwoLists(l1, l2.Next)
	}

	return result
}
