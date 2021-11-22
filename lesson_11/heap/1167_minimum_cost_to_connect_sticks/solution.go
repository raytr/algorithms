package minimum_cost_to_connect_sticks

import "container/heap"

/*
	problem: https://leetcode.com/problems/minimum-cost-to-connect-sticks/

	push sticks to heap


*/
type IntHeap []int

func connectSticks(sticks []int) int {
	if len(sticks) == 1 {
		return 0
	}

	//init heap
	h := IntHeap(sticks)
	heap.Init(&h)

	totalRes := make([]int, 0, len(sticks)-1)
	sum := h.Pop().(int)
	for i := 1; i < len(sticks); i++ {
		val := h.Pop().(int)
		sum += val
		totalRes = append(totalRes, sum)
	}

	sum = 0
	for _, r := range totalRes {
		sum += r
	}
	return sum
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
