package top_k_frequent_elements

import (
	"container/heap"
)

/*
	problem: https://leetcode.com/problems/top-k-frequent-elements

	freqMap map[num]freq
	iterate nums => freqMap => O(n)
	iterate to revert freqMap => freqHeap (val, freq) => n*logn

	for freqHeap.len()	> k
		pop() => k biggest freq  => (n - k)*logn

	for k
		add val to array => k*logn

	complexity: O(nlogn)


	iterage nums
*/

func topKFrequent(nums []int, k int) []int {
	//init heap
	h := initHeap()

	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	for k, v := range freqMap {
		heap.Push(h, FreqVal{
			val:  k,
			freq: v,
		})
	}

	for h.Len() > k {
		heap.Pop(h)
	}

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(FreqVal).val)
	}

	return res
}

//----------------------------------------------------------------------------------

func initHeap() *FreqHeap {
	h := FreqHeap{}
	heap.Init(&h)
	return &h
}

type FreqVal struct {
	freq int
	val  int
}

type FreqHeap []FreqVal

func (h FreqHeap) Len() int           { return len(h) }
func (h FreqHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h FreqHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *FreqHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(FreqVal))
}

func (h *FreqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
