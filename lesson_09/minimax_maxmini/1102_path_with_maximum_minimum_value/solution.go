package path_with_maximum_minimum_value

/*

problem: https://leetcode.com/problems/path-with-maximum-minimum-value/
if a candidate is not satisfied => all of bigger candidates are also not satisfied
if a candidate is satisfied => no need to check any smaller candidates

source <= des => max = source
source >= des => max = des

build dr := [1, 0 , -1, 0]
build dc := [0, 1 , 0, -1]

maxScore = 0
score = source[0,0] = 5
selectedDis := 0
dfs
	while lo <= hi {
		mid = lo + (hi-lo)/2
		if (canDFS(mid)) {
			res = mid
			lo = mid + 1
		} else {
			hi = mid - 1
	}


}


*/

var visited [][]bool
var dr = []int{1, 0, -1, 0}
var dc = []int{0, 1, 0, -1}

func maximumMinimumPath(grid [][]int) int {
	visited = make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}

}

func canDFS(mid int) bool {

}
