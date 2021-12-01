package minimum_cost_to_connect_sticks

import (
	"container/heap"
)

/*
	problem: https://leetcode.com/problems/minimum-cost-to-connect-sticks/
*/

func connectSticks(sticks []int) int {
	h := InitHeaps(sticks)
	sum := 0
	for h.Len() > 1 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)
		sum += a + b
		heap.Push(h, a+b)
	}
	return sum
}

//----------------- Min Heaps -----------------------
type IntHeap []int

func InitHeaps(sticks []int) *IntHeap {
	h := IntHeap(sticks)
	heap.Init(&h)
	return &h
}

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
