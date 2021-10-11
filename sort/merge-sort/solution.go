package main

import "fmt"

//problem: https://leetcode.com/problems/sort-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

var input1 = &ListNode{
	Val: 4,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
			},
		},
	},
}

func main() {
	result := sortList(input1)
	fmt.Println(result)
}

func sortList(head *ListNode) *ListNode {
	// when the linked list just has 1 node
	if head.Next == nil {
		return head
	}

	//merge sort
	//at first, we find the middle node by floyd algorithm
	midNode := getTheMiddleNode(head)

	if midNode == nil {
		return head
	}

	right := midNode.Next
	right = sortList(right)
	midNode.Next = nil
	left := sortList(head)
	return merge2SortedLinkedList(left, right)
}

func getTheMiddleNode(head *ListNode) *ListNode {
	if head.Next == nil || head.Next.Next == nil {
		return head
	}

	fast := head
	slow := head
	for {
		fast = fast.Next.Next

		if fast != nil && fast.Next != nil {
			slow = slow.Next
			continue
		}

		return slow
	}
}

func merge2SortedLinkedList(left, right *ListNode) *ListNode {

	newLL := new(ListNode)

	for left != nil && right != nil {
		if left.Val <= right.Val {
			newLL = &ListNode{
				Val: left.Val,
			}
			left = left.Next
		} else {
			newLL = &ListNode{
				Val: right.Val,
			}
			right = right.Next
		}
		newLL = newLL.Next
	}

	//create node by the rest of l or r linked list
	for left != nil {
		newLL = &ListNode{
			Val: left.Val,
		}
		left = left.Next
		newLL = newLL.Next
	}
	for right != nil {
		newLL = &ListNode{
			Val: right.Val,
		}
		right = right.Next
		newLL = newLL.Next
	}
	return newLL
}
