package k_th_largest_element_in_an_array

import "container/heap"

/*
	problem: https://leetcode.com/problems/kth-largest-element-in-an-array

	complexity: O(n * 2logn)) ~ O(nlogn)
*/

func findKthLargest(nums []int, k int) int {
	//init min heap
	h := initHeap([]int{})

	for _, num := range nums { //O(n)
		heap.Push(h, num) //O(logn)
		if h.Len() > k {
			heap.Pop(h) //O(logN)
		}
	}
	return heap.Pop(h).(int)
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
