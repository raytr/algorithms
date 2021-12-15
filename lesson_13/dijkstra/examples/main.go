package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	//n is number of nodes
	n := 7
	//[src, dest, weight]
	edges := [][]int{{1, 2, 2}, {1, 3, 5}, {2, 3, 6}, {2, 4, 1}, {2, 5, 3}, {4, 5, 4}, {3, 6, 8}, {5, 7, 9}, {6, 7, 7}}
	// start from
	src := 1

	bestWeightTable := Dijkstra(n, edges, src)

	for k, v := range bestWeightTable {
		fmt.Println(fmt.Sprintf("the shortest path from %v to %v is %v", src, k, v))
	}
}

func Dijkstra(n int, edges [][]int, src int) map[int]int {
	adjs := BuildAdjs(n, edges)
	// init min heap and push source with weight is 0 to it
	minHeap := initHeap()
	heap.Push(minHeap, []int{1, 0})

	//init d
	d := make(map[int]int)
	for i := 1; i <= n; i++ {
		d[i] = math.MaxInt32
	}
	d[src] = 0

	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).([]int)
		u := node[0]
		for _, next := range adjs[u] {
			v, vCost := next[0], next[1]
			if d[v] > d[u]+vCost {
				d[v] = d[u] + vCost
				heap.Push(minHeap, []int{v, d[v]})
			}
		}
	}

	return d
}

func BuildAdjs(n int, edges [][]int) map[int][][]int {
	adjs := make(map[int][][]int)
	for i := 1; i <= n; i++ {
		adjs[i] = make([][]int, 0, n-1)
	}
	for _, edge := range edges {
		adjs[edge[0]] = append(adjs[edge[0]], []int{edge[1], edge[2]})
	}
	return adjs
}

//------------------ min heap----------------

//--- heap ---

type MinHeap [][]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func initHeap() *MinHeap {
	h := MinHeap{}
	heap.Init(&h)
	return &h
}

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
