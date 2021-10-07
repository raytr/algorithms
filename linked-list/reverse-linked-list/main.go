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
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

// Recursion (Complexity: O(n))
func reverseListByRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseListByRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
