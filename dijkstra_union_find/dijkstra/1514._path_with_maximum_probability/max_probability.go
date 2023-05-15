package path_with_maximum_probability

import (
	"container/heap"
)

/*
	problem: https://leetcode.com/problems/path-with-maximum-probability/


        build adj list
        build a max heap
        push source to heap

		while with stop condition is maxHeap is empty
			check all current node by ajd list
			if cost = curCost + vCost > max
				push node and cost to heap
*/

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	adjList := buildAdjList(n, edges, succProb) //build adj list
	maxHeap := initMaxHeap()
	maxTable := make(map[float64]float64)
	maxTable[float64(start)] = 1

	heap.Push(maxHeap, []float64{float64(start), succProb[0]})

	for maxHeap.Len() > 0 {
		node := heap.Pop(maxHeap).([]float64)
		u := int(node[0])
		if u == end {
			return maxTable[float64(u)]
		}
		for _, e := range adjList[u] {
			v := int(e[0])
			vCost := e[1]
			if maxTable[float64(v)] < maxTable[float64(u)]*vCost {
				maxTable[float64(v)] = maxTable[float64(u)] * vCost
				heap.Push(maxHeap, []float64{float64(v), maxTable[float64(v)]})
			}
		}
	}
	return 0
}

func buildAdjList(n int, edges [][]int, succProb []float64) map[int][][]float64 {
	adjMap := make(map[int][][]float64)

	for i, edge := range edges {
		adjMap[edge[0]] = append(adjMap[edge[0]], []float64{float64(edge[1]), succProb[i]})
		adjMap[edge[1]] = append(adjMap[edge[1]], []float64{float64(edge[0]), succProb[i]})
	}

	return adjMap
}

// -------- build max heap
type MaxHeap [][]float64

func initMaxHeap() *MaxHeap {
	h := MaxHeap{}
	heap.Init(&h)
	return &h
}

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]float64))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
