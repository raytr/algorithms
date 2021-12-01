package connecting_cities_with_minimum_cost

import "sort"

/*

	problem: https://leetcode.com/problems/connecting-cities-with-minimum-cost/


	   sort by cost
	   union, if success -> sum += cost; count++

	   if count < len(city) - 1 => return -1


	time complexity
*/

type DisjointSet struct {
	size         int
	parent, rank []int
}

func minimumCost(n int, connections [][]int) int {
	//int disjoint set
	djs := InitDisjointSet(n)

	connectedCount, sumCost := 0, 0

	//sort
	sort.Slice(connections, func(a, b int) bool {
		return connections[a][2] < connections[b][2]
	})

	for _, connection := range connections {
		if djs.union(connection[0], connection[1]) {
			connectedCount++
			sumCost += connection[2]
		}
	}

	if connectedCount < n-1 {
		return -1
	}

	return sumCost
}

func InitDisjointSet(n int) DisjointSet {
	djs := DisjointSet{
		size:   n,
		parent: make([]int, n+1, n+1),
		rank:   make([]int, n+1, n+1),
	}
	for i := 1; i <= n; i++ {
		djs.parent[i] = i
	}
	return djs
}

func (this *DisjointSet) union(x, y int) bool {
	rootX, rootY := this.find(x), this.find(y)
	if rootX == rootY {
		return false
	}

	if this.rank[rootX] > this.rank[rootY] {
		this.parent[rootY] = rootX
	} else if this.rank[rootY] > this.rank[rootX] {
		this.parent[rootX] = rootY
	} else {
		this.parent[rootY] = rootX
	}
	return true
}

func (this *DisjointSet) find(x int) int {
	if this.parent[x] == x {
		return this.parent[x]
	} else {
		return this.find(this.parent[x])
	}
}
