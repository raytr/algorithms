package main

//problem: https://leetcode.com/problems/middle-of-the-linked-list

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//Complexity: O(n)
func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
