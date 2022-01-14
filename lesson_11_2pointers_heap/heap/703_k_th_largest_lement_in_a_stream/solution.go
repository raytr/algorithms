package k_th_largest_lement_in_a_stream

import (
	"container/heap"
)

/*
	problem: https://leetcode.com/problems/kth-largest-element-in-a-stream
	complexity: O(logn)
*/

type KthLargest struct {
	h *IntHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	// init heap
	h := IntHeap(nums)
	heap.Init(&h)
	return KthLargest{
		h: &h,
		k: k,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.h, val)
	for this.h.Len() > this.k {
		heap.Pop(this.h)
	}
	res := heap.Pop(this.h).(int)
	heap.Push(this.h, res)
	return res
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

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
