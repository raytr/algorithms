package min_cost_to_connect_all_points

import (
	"math"
	"sort"
)

/*

	problem: https://leetcode.com/problems/min-cost-to-connect-all-points/


        build list of edges between nodes follow format
            [node1, node2, cost]
		we will index node by formula :
		idx := (x+1000000)*1000000+y+1000000
				(+1000000 to solve negative number)

        build union find
        sort all of edges that we just build
        go over of edges, with each, union(edge[1], edge[2])
        if union() == true
            cost += edge[2]

        return cost


		complexity:
*/

type DisjointSet struct {
	size         int
	parent, rank map[int64]int64
}

func minCostConnectPoints(points [][]int) int {
	edges := make([][]int64, 0)
	uniqueCdnVal := make(map[int64]bool)
	cost := int64(0)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			// cost = |xi - xj| + |yi - yj|
			cost := math.Abs(float64(points[i][0]-points[j][0])) + math.Abs(float64(points[i][1]-points[j][1]))
			//index point
			node1 := int64((points[i][0]+1000000)*1000000 + points[i][1] + 1000000)
			node2 := int64((points[j][0]+1000000)*1000000 + points[j][1] + 1000000)
			uniqueCdnVal[node1] = true
			uniqueCdnVal[node2] = true
			edges = append(edges, []int64{node1, node2, int64(cost)})
		}
	}

	//sort the edges we just build
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	//build djs
	djs := InitDisjointSet(uniqueCdnVal)
	for _, e := range edges {
		if djs.union(e[0], e[1]) {
			cost += e[2]
		}
	}
	return int(cost)
}

func InitDisjointSet(m map[int64]bool) DisjointSet {
	n := len(m)
	djs := DisjointSet{
		size:   n,
		parent: make(map[int64]int64),
		rank:   make(map[int64]int64),
	}
	for num, _ := range m {
		djs.parent[num] = num
		djs.rank[num] = 0
	}
	return djs
}

func (this *DisjointSet) union(x, y int64) bool {
	rootX, rootY := this.find(x), this.find(y)
	if rootX == rootY {
		return false
	}

	if this.rank[rootY] > this.rank[rootX] {
		this.parent[rootX] = rootY
	} else if this.rank[rootX] > this.rank[rootY] {
		this.parent[rootY] = rootX
	} else {
		this.parent[rootY] = rootX
		this.rank[rootX]++
	}
	return true
}

func (this *DisjointSet) find(x int64) int64 {
	if x == this.parent[x] {
		return this.parent[x]
	} else {
		return this.find(this.parent[x])
	}
}
