package maximum_score_from_removing_stones

import (
	"container/heap"
)

/*
	problem: https://leetcode.com/problems/maximum-score-from-removing-stones

	   use heap
	   at this problem, we always get 2 stones in 2 piles has most stones

	   while zeroPiles <= 1
		   use heap, pop 1 first => 2 rest piles will are 2 piles has most stones
		   pop 2 rest piles and with each val we -1 => if val == 0, zeroPiles++
		   push all val to heap

	complexity: O(logn)
*/

func maximumScore(a int, b int, c int) int {
	// init heap
	h := initHeap(a, b, c)

	score, emptyPiles := 0, 0
	for emptyPiles <= 1 {
		a = heap.Pop(h).(int) // O(logn)
		b = heap.Pop(h).(int) // O(logn)
		c = heap.Pop(h).(int) // O(logn)

		b -= 1
		if b == 0 {
			emptyPiles++
		}
		c -= 1
		if c == 0 {
			emptyPiles++
		}

		// push all of val to heap
		heap.Push(h, a)
		heap.Push(h, b) // O(logn)
		heap.Push(h, c) // O(logn)
		score++
	}
	return score
}

func initHeap(a, b, c int) *IntHeap {
	h := IntHeap{a, b, c}
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
