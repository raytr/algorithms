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
//if  l1.Val <=  l2.Val {
//result = &ListNode{
//Val:  l1.Val,
//}
//l1 =  l1.Next
//}else {
//result = &ListNode{
//Val:  l1.Val,
//}
//l2 =  l2.Next
//}
//
//for  l1 == nil ||  l2 == nil {
//if  l1.Val <=  l2.Val {
//result.Next =  l1.Val
//l1 =  l1.Next
//} else {
//result.Next =  l2.Val
//l2 =  l2.Next
//}
//result = result.Next
//}
//
////create node by the rest of l or r linked list
//for  l1 != nil {
//result.Next =  l1.Val
//l1 =  l1.Next
//result = result.Next
//}
//for  l2 != nil {
//result.Next =  l2.Val
//l2 =  l2.Next
//result = result.Next
//}
//return result
