package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	Src    string
	Des    string
	Weight float64
}

type Graph struct {
	Nodes []string
	Edges []Edge
}

func main() {

	g := Graph{
		Nodes: []string{"a", "b", "c", "d", "e", "f", "g"},
	}
	g.AddEdge("a", "b", 2)
	g.AddEdge("a", "c", 5)
	g.AddEdge("b", "c", 6)
	g.AddEdge("b", "d", 1)
	g.AddEdge("b", "e", 3)
	g.AddEdge("d", "e", 4)
	g.AddEdge("c", "f", 8)
	g.AddEdge("e", "g", 9)
	g.AddEdge("f", "g", 7)

	adjs := g.BuildAdjs()

	fmt.Println(g.dijkstra("a"))
}

//	priority_queue<ii, vector<ii>, greater<ii>> q;
//// min heap, each element in heap is a pair (d[v], v)
//int d[N+1];
//for(int i = 1; i <= N; i++) d[i] = INT_MAX;
//d[Source] = 0; // u != source => d[u] = oo => no need add u to Q
//q.push(ii(0, Source));
//// adj[u] = { (v1, w1), (v2, w2), … }
//while (q.empty() == false) {
//int u = q.top().second; // add u to S
//q.pop(); // n times => O(n*logn)
//for(auto pair: adj[u]) { // all v in adj[u]
//int v = pair.first; // node
//int w = pair.second; // weight
//if (d[v] > d[u] + w) {
//d[v] = d[u] + w;
//prev[v] = u;
//q.push(ii(d[v], v)); // O(logn * m)
//}
//}
//}

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

func (g *Graph) AddEdge(s, d string, w float64) {
	edge := Edge{
		Src:    s,
		Des:    d,
		Weight: w,
	}

	g.Edges = append(g.Edges, edge)
}

func (g *Graph) BuildAdjs() map[string][]Edge {
	adjs = make(map[string][]Edge)
	for _, edge := range g.Edges {
		g.Adjs[edge.Src] = append(g.Adjs[edge.Src], edge)
	}
	return adjs
}

//------------------ min heap----------------
type MinHeap []Edge

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
