package main

import (
	"container/heap"
	"math"
)

type Node struct {
	Name   string
	Weight float64
}

func main() {

}

func (g *Graph) dijkstra(startNode string) string {
	// init heap
	propertiesQueue := initHeap(g)

	//init d
	d := make(map[string]float64)
	for _, n := range g.Nodes {
		d[n] = math.Inf(1)
	}
	d[startNode] = 0

	for propertiesQueue.Len() > 0 {
		node := heap.Pop(propertiesQueue).(Edge)
		for _, v := range g.Adjs[node.Src] {
			if d[v.Des] > d[v.Src]+v.Weight {
				d[v.Des] = d[v.R]
			}
		}
	}

	return ""
}

//------------------ min heap----------------
type MinHeap []Node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Weight < h[j].Weight }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func initHeap(g *Graph) *MinHeap {
	h := MinHeap(g.Edges)
	heap.Init(&h)
	return &h
}

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Edge))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
