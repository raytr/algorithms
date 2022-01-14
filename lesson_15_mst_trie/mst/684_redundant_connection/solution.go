package redundant_connection

import "math"

/*

	problem: https://leetcode.com/problems/redundant-connection/

*/

type DisjointSet struct {
	size         int
	parent, rank []int
}

func findRedundantConnection(edges [][]int) []int {
	//	//get max number
	//	max := 0
	//	for _, e := range edges {
	//		max = int(math.Max(float64(e[0]), float64(max)))
	//		max = int(math.Max(float64(e[1]), float64(max)))
	//	}
	//
	//	djs := initDisjointSet(max)
	//
	//	//find the root that has max frequency
	//	freqCount := make(map[int]int)
	//	for _, e := range edges {
	//		if !djs.union(e[0], e[1]) {
	//			return e
	//		}
	//	}
	//	return nil
	//}
	//
	//func initDisjointSet(n int) DisjointSet {
	//	this := DisjointSet{
	//		size:   n,
	//		parent: make([]int, n+1, n+1),
	//		rank:   make([]int, n+1, n+1),
	//	}
	//	for i := 1; i <= n; i++ {
	//		this.parent[i] = i
	//	}
	//	return this
	//}
	//
	//func (this *DisjointSet) union(x,y int) bool {
	//	rootX, rootY := this.parent[x], this.parent[y]
	//	if rootX == rootY {
	//		return false
	//	}
	//	if this.rank[rootX] > this.rank[rootY] {
	//		this.parent[rootY] = rootX
	//	} else if this.rank[rootX] < this.rank[rootY] {
	//		this.parent[rootX] = rootY
	//	} else {
	//		this.parent[rootY] = rootX
	//		this.rank[rootX]++
	//	}
	//	return true
	//}
	//
	//func (this *DisjointSet)find(x int) int {
	//	if x == this.parent[x] {
	//		return this.parent[x]
	//	} else {
	//		return this.find(this.parent[x])
	//	}
}
