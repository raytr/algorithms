package palindrome_linked_list

/*
	problem: https://leetcode.com/problems/palindrome-linked-list/

    avoid to use extra memmory

    step 1: find the center node
        use 2 pointers, fast and slow
        when fash move 2 note, slow move 1 node. => when fast.Next  == nil => centre node is slow

   step 2: revert ll half behind
   step 3: compare util meet a centre node

	Time complexity: O(n)
	Space complexity: O(1)
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	res := true
	p1 := head
	p2 := endOfFirstHaft(head)
	p2 = reverseList(p2)
	for p2 != nil {
		if p2.Val != p1.Val {
			res = false
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return res
}

func endOfFirstHaft(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
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
