package cheapest_flights_within_k_stops

import (
	"container/heap"
	"math"
)

/*
	problem: https://leetcode.com/problems/cheapest-flights-within-k-stops/

        build adj map based on k stop
         adjM[0] = [[1,1], [2,1]]
         adjM[1] = [[2,0]]

        build the priorityQueue is a min heap
        betterWayMap is a map with key neatest node and value is array [cost, stop]

        while with stop condition is queue is empty
            node = heap.pop
			cur = heap[0]
			cost = heap[1]
			stop = heap[2]
			if  stop > 0
				for next range adjMap[u]
					v = next[0]
					vCost = next[1]
					if uCost > vCost+w => update betterWayMap with better cost, update better stop => queue.push(v, cost + nCost, k-1)
					else uStop < vStop => update betterWayMap with better stop => => queue.push(v, cost + nCost, k-1)

	complexity:
		n is number of node
		m is number of edge
*/

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	adjMap := make(map[int][][]int)
	for _, flight := range flights {
		if _, exist := adjMap[flight[0]]; !exist {
			adjMap[flight[0]] = make([][]int, 0, len(flights))
		}
		adjMap[flight[0]] = append(adjMap[flight[0]], []int{flight[1], flight[2]})
	}

	newBetterWayMap := make(map[int][]int) // get better way via 2 things: cost and stop
	for i := 0; i < n; i++ {
		newBetterWayMap[i] = []int{math.MaxInt32, 0}
	}
	newBetterWayMap[src] = []int{0, k + 1}

	//minHeap based on cost of distance from source
	minHeap := initHeap()
	heap.Push(minHeap, []int{src, 0, k + 1})

	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).([]int) //O(nLogn)
		u, uCost, uStop := node[0], node[1], node[2]
		if u == dst {
			return uCost
		}
		if uStop > 0 {
			for _, next := range adjMap[u] {
				v, vCost := next[0], next[1]
				vStop := uStop - 1
				position := newBetterWayMap[v]
				if position[0] > uCost+vCost { //check to get better cost
					newBetterWayMap[v][0] = uCost + vCost
					newBetterWayMap[v][1] = vStop
					heap.Push(minHeap, []int{v, vCost + uCost, uStop - 1})
				} else if newBetterWayMap[v][1] < vStop { //check to get better stop
					heap.Push(minHeap, []int{v, vCost + uCost, uStop - 1})
				}
			}
		}
	}
	return -1
}

//--- heap ---

type FlightsHeap [][]int

func (h FlightsHeap) Len() int           { return len(h) }
func (h FlightsHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h FlightsHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func initHeap() *FlightsHeap {
	h := FlightsHeap{}
	heap.Init(&h)
	return &h
}

func (h *FlightsHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *FlightsHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
