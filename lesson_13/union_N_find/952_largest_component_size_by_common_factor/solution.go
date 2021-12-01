package largest_component_size_by_common_factor

/*

	problem: https://leetcode952.com/problems/largest-component-size-by-common-factor/


        union find sulution:
        with number A, we find its factors; merge a and it's factors

        after build union-find structure,
            - with each number, find its root and count[root]++
            - find the max root


        return the max root


	complexity:
*/

type DisjointSet struct {
	size         int
	rank, parent []int
}

func largestComponentSize(nums []int) int {
	//find max
	max := 0
	for _, n := range nums {
		max = Max(max, n)
	}
	//build union findD
	djs := initDisjointSet(max)

	//we iterate through each number. For each number , we iterate from 2 to sqrt(num) to find all factor
	for _, num := range nums {
		for factor := 2; factor*factor <= num; factor++ {
			if num%factor == 0 {
				// for each factor of num, we union it
				djs.union(num, factor)
				djs.union(num, num/factor)
			}
		}
	}

	maxCount := 0
	rootCount := make(map[int]int)
	for _, num := range nums {
		root := djs.find(num)
		rootCount[root]++
		maxCount = Max(maxCount, rootCount[root])
	}
	return maxCount
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

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
