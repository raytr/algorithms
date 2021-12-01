package graph_connectivity_with_threshold

/*
	problem: https://leetcode.com/problems/graph-connectivity-with-threshold/
        build union-set follow n

        we iterate through each number from threshold + 1,
        with each number, we iterate from threshold + 1
            if num odd divisor
                union(num, divisor)
                union(num, num/divisor)

        after build union-set we have many roots
        we check pair by pair, if 2 number belong this pair has common root
            => add true to res, otherwise add false
*/

type DisjointSet struct {
	size         int
	rank, parent []int
}

func areConnected(n int, threshold int, queries [][]int) []bool {
	//build union-set
	djs := initDisjointSet(n)
	for i := threshold + 1; i <= n; i++ {
		for devisor := threshold + 1; devisor <= i; devisor++ {
			if i%devisor == 0 {
				djs.union(i, devisor)
				if i/devisor > threshold {
					djs.union(i, i/devisor)
				}
			}
		}
	}

	res := make([]bool, 0, len(queries))
	for _, q := range queries {
		if djs.find(q[0]) == djs.find((q[1])) {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}

func (this *DisjointSet) union(x, y int) {
	rootX, rootY := this.find(x), this.find(y)

	if rootX == rootY {
		return
	}

	if this.rank[rootX] > this.rank[rootY] {
		this.parent[rootY] = rootX
	} else if this.rank[rootY] > this.rank[rootX] {
		this.parent[rootX] = rootY
	} else {
		this.parent[rootY] = rootX
		this.rank[rootX]++
	}
}

func (this *DisjointSet) find(x int) int {
	if x == this.parent[x] {
		return this.parent[x]
	} else {
		return this.find(this.parent[x])
	}
}

func initDisjointSet(n int) DisjointSet {
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
