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
	rank, parent map[int]int
}

func largestComponentSize(nums []int) int {
	uniqueMap := make(map[int]bool)
	//factorOf := make(map[int][]int)
	for _, num := range nums {
		m := getFactor(num)
		for k, _ := range m {
			uniqueMap[k] = true
		}
		uniqueMap[num] = true
	}

	//build union findD
	djs := initDisjointSet(uniqueMap)
	uniqueNumArr := make([]int, 0, len(uniqueMap))
	for k, _ := range uniqueMap {
		uniqueNumArr = append(uniqueNumArr, k)
	}

	// merge all factors if a to each other
	arr := make([]int, 0, len(uniqueMap))
	for num, _ := range uniqueMap {
		arr = append(arr, num)
	}
	for i := 0; i < len(arr)-1; i++ {
		djs.union(arr[i], arr[i+1])
	}

	// find the root with max frequently
	rootCounts := make(map[int]int)
	for _, n := range nums {
		rootCounts[djs.find(n)]++
	}

	maxCount := make([]int, 2, 2) // idx 0th is root, idx 1th is freq
	for k, v := range rootCounts {
		if v > maxCount[1] {
			maxCount = []int{k, v}
		}
	}

	return maxCount[1]
}

func getFactor(num int) map[int]bool {
	primeMap := make(map[int]bool)
	for i := 2; i*i <= num; i++ {
		for num%i == 0 {
			primeMap[i] = true
			num /= i
		}
	}

	if num != 1 {
		primeMap[num] = true
	}

	return primeMap
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

func initDisjointSet(uniqueMap map[int]bool) DisjointSet {
	n := len(uniqueMap)
	djs := DisjointSet{
		size:   n,
		parent: make(map[int]int),
		rank:   make(map[int]int),
	}

	for k, _ := range uniqueMap {
		djs.parent[k] = k
		djs.rank[k] = 0
	}

	return djs

}
