package merge_k_sorted_lists

import "container/heap"

/*
    problem: https://leetcode.com/problems/merge-k-sorted-lists/
   bruforce:
       iterate all list and push all vals to array O(n)
        init heap follow arr this array

        pop heap until heap is empty O(nLogn)

    complexity: O(nLogN)

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	arr := make([]int, 0)
	for _, list := range lists {
		for {
			if list == nil {
				break
			}
			arr = append(arr, list.Val)
			list = list.Next
		}
	}

	h := initHeap(arr)
	res := setList(h)

	return res
}

func setList(h *IntHeap) *ListNode {
	if h.Len() == 0 {
		return nil
	}
	list := &ListNode{
		Val:  heap.Pop(h).(int),
		Next: setList(h),
	}
	return list
}

// ------ heap --------

func initHeap(arr []int) *IntHeap {
	h := IntHeap(arr)
	heap.Init(&h)
	return &h
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
