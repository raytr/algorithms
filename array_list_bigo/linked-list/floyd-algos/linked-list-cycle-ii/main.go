package main

//problem: https://leetcode.com/problems/linked-list-cycle-ii

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//Complexity: O(n)
func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast != nil && slow != nil && fast == slow {
			return true
		}
	}
	return false
}
