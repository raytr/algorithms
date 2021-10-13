package main

//problem: https://leetcode.com/problems/reverse-linked-list

//Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

// Iterative  (Complexity: O(n))
func reverseListByIterative(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		theNext := cur.Next
		cur.Next = prev
		prev = cur
		cur = theNext
	}

	return prev
}

// Recursion (Complexity: O(n))
func reverseListByRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseListByRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
