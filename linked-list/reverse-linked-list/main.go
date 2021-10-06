package reverse_linked_list

//problem: https://leetcode.com/problems/reverse-linked-list

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//Complexity: O(n)
func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}
