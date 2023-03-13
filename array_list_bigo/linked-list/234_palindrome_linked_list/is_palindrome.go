package palindrome_linked_list

/*
	problem: https://leetcode.com/problems/palindrome-linked-list/

    avoid to use extra memmory

	clarify:
    what is the range of the number of nodes in this linked list
    what is range of node's value

    avoid to use extra memmory

    step 1: find the center node
        use 2 pointers, fast and slow
        when fast move 2 note, slow move 1 node. => when fast.Next  == nil => centre node is slow

   step 2: revert ll half behind
   step 3: compare util meet a centre node

    Time complexity: O(n)
    Space complexity: O(1)
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*


 */

func isPalindrome(head *ListNode) bool {
	//eage case
	if head.Next == nil {
		return true
	}

	centerNode := findCenterNode(head)

	//revert linked list from half to end
	revertedLL := revertLinkedList(centerNode)

	//compare node by node until end of revertedLL
	for revertedLL != nil {
		if revertedLL.Val != head.Val {
			return false
		}
		head = head.Next
		revertedLL = revertedLL.Next
	}
	return true
}

func findCenterNode(head *ListNode) *ListNode {
	fast := head.Next.Next
	slow := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func revertLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := revertLinkedList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
