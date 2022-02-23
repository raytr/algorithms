package number_of_islands_ii

/*
	problem: https://leetcode.com/problems/number-of-islands-ii/

        at first, index for all of cell by fomular: index = row*m + col

        islandTh = 0
        go over all of positions
            islandTh++
            with each position
                look 4 direction, if can union => islandTh--
                res.push(islandTh)

*/

type disjointSet struct {
	size         int
	parent, rank []int
}

func numIslands2(m int, n int, positions [][]int) []int {

	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}
	islandTh := 0
	djs := initDisjointSet(m * n)
	res := make([]int, 0, len(positions))
	grid := make([][]int, m, m*n)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n, n)
	}

	for _, p := range positions {
		curIdx := p[0]*n + p[1]
		if grid[p[0]][p[1]] == 0 {
			islandTh++
			grid[p[0]][p[1]] = 1
		}
		for i := 0; i < 4; i++ {
			u := p[0] + dr[i]
			v := p[1] + dc[i]
			nearestIslandIdx := u*n + v
			if u >= 0 && u < m && v >= 0 && v < n && grid[u][v] == 1 && djs.union(nearestIslandIdx, curIdx) {
				islandTh--
			}
		}
		res = append(res, islandTh)
	}
	return res
}

func initDisjointSet(n int) disjointSet {
	djs := disjointSet{
		size:   n,
		parent: make([]int, n, n),
		rank:   make([]int, n, n),
	}
	for i := 0; i < n; i++ {
		djs.parent[i] = i
	}
	return djs
}

func (this *disjointSet) union(x, y int) bool {
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
		this.rank[rootX]++
	}
	return true
}

func (this *disjointSet) find(x int) int {
	if this.parent[x] == x {
		return this.parent[x]
	} else {
		return this.find(this.parent[x])
	}
}
