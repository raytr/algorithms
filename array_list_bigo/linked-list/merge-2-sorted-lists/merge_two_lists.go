package merge_2_sorted_lists

// problem: https://leetcode.com/problems/merge-two-sorted-lists/
type ListNode struct {
	Val  int
	Next *ListNode
}

// complexity: O[len(l1) + len(l2)]
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	result := new(ListNode)
	if l1.Val <= l2.Val {
		result.Val = l1.Val
		result.Next = mergeTwoLists(l1.Next, l2)
	} else {
		result.Val = l2.Val
		result.Next = mergeTwoLists(l1, l2.Next)
	}

	return result
}

//// create a first node for new LL
//if  l1.val <=  l2.val {
//result = &ListNode{
//val:  l1.val,
//}
//l1 =  l1.next
//}else {
//result = &ListNode{
//val:  l1.val,
//}
//l2 =  l2.next
//}
//
//for  l1 == nil ||  l2 == nil {
//if  l1.val <=  l2.val {
//result.next =  l1.val
//l1 =  l1.next
//} else {
//result.next =  l2.val
//l2 =  l2.next
//}
//result = result.next
//}
//
////create node by the rest of l or r linked list
//for  l1 != nil {
//result.next =  l1.val
//l1 =  l1.next
//result = result.next
//}
//for  l2 != nil {
//result.next =  l2.val
//l2 =  l2.next
//result = result.next
//}
//return result
