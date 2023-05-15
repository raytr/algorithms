package remove_zero_sum_consecutive_nodes_from_linked_list

/*
 https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/
 time complexity: O(n)
 used map to store -> space complexity is : O(n)
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	//a dummy node
	dummy := &ListNode{Val: 0, Next: head}
	sum := 0
	//create a map for now sum and point
	hashMap := make(map[int]*ListNode)
	for i := dummy; i != nil; i = i.Next {
		sum += i.Val
		//if have same sum ,cover odl value
		hashMap[sum] = i
	}
	sum = 0
	for i := dummy; i != nil; i = i.Next {
		sum += i.Val
		i.Next = hashMap[sum].Next
	}
	return dummy.Next
}
