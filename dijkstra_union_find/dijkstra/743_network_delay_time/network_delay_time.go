package network_delay_time

import (
	"container/heap"
	"math"
)

/*
	problem: https://leetcode.com/problems/network-delay-time/

	to calculate the minimum time it takes for all the n nodes to receive the signal,
	we need to find the maximum path
*/

func networkDelayTime(times [][]int, n int, k int) int {
	adjM := buildAdjs(times)
	visitedMap := make(map[int]bool)
	visitedMap[k] = true
	max := 0

	//init queue
	priorityQ := initHeap()
	heap.Push(priorityQ, []int{k, 0})

	//init table
	table := make(map[int]int)
	for i := 1; i <= n; i++ {
		table[i] = math.MaxInt32
	}
	table[k] = 0

	for priorityQ.Len() > 0 {
		u := heap.Pop(priorityQ).([]int)[0]

		for _, pair := range adjM[u] {
			v := pair[0]
			w := pair[1]
			visitedMap[v] = true
			if table[v] > table[u]+w {
				table[v] = table[u] + w
				heap.Push(priorityQ, []int{v, table[v]})
			}
		}
	}

	//check the whole of nodes was visited
	if len(visitedMap) < n {
		return -1
	}

	for _, v := range table {
		max = getMax(max, v)
	}

	return max
}

func buildAdjs(time [][]int) map[int][][]int {
	adjM := make(map[int][][]int)
	for _, t := range time {
		if _, exist := adjM[t[0]]; !exist {
			adjM[t[0]] = make([][]int, 0)
		}
		adjM[t[0]] = append(adjM[t[0]], []int{t[1], t[2]})
	}

	return adjM
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//--- heap ---

type TimeHeap [][]int

func (h TimeHeap) Len() int           { return len(h) }
func (h TimeHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h TimeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func initHeap() *TimeHeap {
	h := TimeHeap{}
	heap.Init(&h)
	return &h
}

func (h *TimeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *TimeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
