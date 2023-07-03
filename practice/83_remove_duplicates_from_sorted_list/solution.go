package remove_duplicates_from_sorted_list

import "fmt"

/*
	problem: https://leetcode.com/problems/remove-duplicates-from-sorted-list/


   1 - 1 - 1- 1 - 2
   cur = head
   while cur != nil
       while cur.val == nextNode.val
           nextNode.val == cur.val => cur.next = nextNode.next
       cur = nextNode

   return

   time complexity : O(n)
   space : O(1)
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		fmt.Println(cur.Val)
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}
