//package minimum_cost_to_reach_city_with_discounts
package main

import (
	"fmt"
	"sort"
)

/*
	problem: https://leetcode.com/problems/minimum-cost-to-reach-city-with-discounts/

        use spt to solve

        add a variable represent for discount, with source = k, otherwise node is 0 => O(n)
        with each node, we add a node with cost is this node's value / 2,
			- mark the discounts = -1 => O(n)
			- with node is m + index

		sort new array

        we have a rule, when connect 2 cities,
            => the cost is city1's discounts + city2's discounts
            if discounts is < 0 can't union
            ortherwise union
            when union success => cost = a[2] + b[2]
            the complexity to union is O(nlogn)

        if find(src) != find(des) => return -1

        return the cost

	complexity : O(nlogN)
*/

func main() {
	fmt.Println(minimumCost(5, [][]int{{0, 1, 4}, {2, 1, 3}, {1, 4, 11}, {3, 2, 3}, {3, 4, 2}}, 1))
}

type DisjointSet struct {
	size         int
	parent, rank []int
}

func minimumCost(n int, highways [][]int, discounts int) int {

	//for i := 0; i < n; i++ {
	//	highways[i] = append(highways[i], 0)
	//}
	//highways[0][3] = discounts
	//
	////add fake cities with cost is cost of real city / 2 and index it with fomular := n + i
	//// for example : [0,2,3,0] => [0, 2+n, 3/2, -1]
	//for i := 1; i < n; i++ {
	//	highways = append(highways, []int{highways[i][0], highways[i][1] + n, highways[i][2] / 2, -1})
	//}

	//sort by cost
	sort.Slice(highways, func(i, j int) bool {
		return highways[i][2] < highways[j][2]
	})

	//init union set
	djs := initDisjoinSet(n*2 - 1)

	cost := 0
	for _, hw := range highways {
		if djs.union(hw[0], hw[1]) {
			cost += hw[2]
		}
	}

	if djs.find(0) != djs.find(n-1) {
		return -1
	}

	return cost
}

//----------- Union Find ------------

func initDisjoinSet(n int) DisjointSet {
	djs := DisjointSet{
		size:   n,
		parent: make([]int, n, n),
		rank:   make([]int, n, n),
	}
	for i := 0; i < n; i++ {
		djs.parent[i] = i
	}
	return djs
}

func (this *DisjointSet) union(x, y int) bool {
	rootX, rootY := this.find(x), this.find(y)
	if rootX == rootY {
		return false
	}

	if this.rank[rootY] > this.rank[rootX] {
		this.parent[rootX] = rootY
	} else if this.rank[rootY] < this.rank[rootX] {
		this.parent[rootY] = rootX
	} else {
		this.parent[rootY] = rootX
		this.rank[rootX]++
	}
	return true
}

func (this *DisjointSet) find(x int) int {
	if x == this.parent[x] {
		return this.parent[x]
	} else {
		this.parent[x] = this.find(this.parent[x])
		return this.parent[x]
	}
}
