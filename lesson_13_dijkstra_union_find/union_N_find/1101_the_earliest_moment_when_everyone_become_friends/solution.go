package _101_the_earliest_moment_when_everyone_become_friends

import "sort"

/*

	problem :https://leetcode.com/problems/the-earliest-moment-when-everyone-become-friends/

	Use an additional array size[0..n] to store the size of each tree.
	size[i] = 1

	void union(int u, int v) {
		int rootU = find(u);
		int rootV = find(v);
		if (rootU < rootV) {
			parent[rootU] = rootV;
			size[rootU] += size[rootV]
			if size[rootU] == n: return
		} else {
			parent[rootV] = rootU;
			size[rootV] += size[rootU]
			if size[rootV] == n: return
				if (rank[rootU] == rank[rootV]) {
					rank[rootU]++;
				}
			}
		}


	sort logs based on timestemp
	create disjointsSet with parrent[x] == x; range all of them == 0
	go over all of logs
		if union()
			n--;
		if n == 1 => return log[i].timestamp

	`return -1
*/

type DisjointSet struct {
	size         int
	rank, parent []int
}

func earliestAcq(logs [][]int, n int) int {
	//sort to make sure the last union at the latest timestamp
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	djs := initDisjointSet(n)

	for i := 0; i < len(logs); i++ {
		if djs.union(logs[i][1], logs[i][2]) {
			n--
		}

		if n == 1 {
			return logs[i][0]
		}
	}

	return -1
}

func (this *DisjointSet) union(x, y int) bool {
	rootX := this.find(x)
	rootY := this.find(y)
	if rootX == rootY {
		return false
	}

	if this.rank[rootX] > this.rank[rootY] {
		this.parent[rootY] = rootX
	} else if this.rank[rootX] < this.rank[rootY] {
		this.parent[rootX] = rootY
	} else {
		this.parent[rootY] = rootX
		this.rank[rootX]++
	}
	return true
}

func initDisjointSet(n int) DisjointSet {
	djs := DisjointSet{
		size:   n,
		parent: make([]int, n, n),
		rank:   make([]int, n, n),
	}
	for i := 0; i < n; i++ {
		djs.parent[i] = i
		djs.rank[i] = 0
	}
	return djs
}

func (this *DisjointSet) find(x int) int {
	if x == this.parent[x] {
		return this.parent[x]
	} else {
		return this.find(this.parent[x])
	}
}
