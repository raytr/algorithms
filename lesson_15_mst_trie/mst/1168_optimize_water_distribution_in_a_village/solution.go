package optimize_water_distribution_in_a_village

import "sort"

/*
	problem: https://leetcode.com/problems/optimize-water-distribution-in-a-village/

        image we have a source, and this sorce connect to all of houses.
        the cost to build a well is the cost to pipe in water from this source to all of the houses

        => the problem to change to minimum spanning tree problem

        for idx, well of range wells
            pipes.push(0, idx + 1, well)
        sort pipes

        build a union set
        go over the pipes ,
        if union (pipe[0], pipe1) == true
            cost++

        return cost


	time complexity :
*/

type DisjointSet struct {
	size         int
	parent, rank []int
}

func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	//build union
	djs := InitDisjointSet(n)
	cost := 0

	for idx, well := range wells {
		pipes = append(pipes, []int{0, idx + 1, well})
	}
	sort.Slice(pipes, func(i, j int) bool {
		return pipes[i][2] < pipes[j][2]
	})

	for _, pipe := range pipes {
		if djs.union(pipe[0], pipe[1]) {
			cost += pipe[2]
		}
	}
	return cost
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
