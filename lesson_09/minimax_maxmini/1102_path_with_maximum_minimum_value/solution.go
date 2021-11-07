package path_with_maximum_minimum_value

/*

problem: https://leetcode.com/problems/path-with-maximum-minimum-value/
if a candidate is not satisfied => all of bigger candidates are also not satisfied
if a candidate is satisfied => no need to check any smaller candidates

build dr := [1, 0 , -1, 0]
build dc := [0, 1 , 0, -1]



Lo = 0; hi = 200; M = 100
in part 1: Meet 5 => mid can not is score because score is min in path
			=> return false => hi = mid - 1 (because mid should be smaller)
..............
Hi = 4; Lo = 0; M = 2
	Meet 1 => mid satisfied because mid can be score => res = mid,  lo = mid + 1 (try to find smaller mid)

*/

var visited [][]bool
var dr = []int{1, 0, -1, 0}
var dc = []int{0, 1, 0, -1}

func maximumMinimumPath(grid [][]int) int {
	lo, hi := 0, int(10e9)
	res := 0

	for lo <= hi {
		visited = make([][]bool, len(grid)+1)
		for i := 0; i < len(grid); i++ {
			visited[i] = make([]bool, len(grid[i])+1)
		}
		visited[0][0] = true
		mid := lo + (hi-lo)/2
		if mid <= grid[0][0] {
			if canDFS(grid, mid, 0, 0) {
				res = mid
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else {
			hi = mid - 1
		}
	}
	return res
}

func canDFS(grid [][]int, mid, u, v int) bool {
	visited[u][v] = true
	if u == len(grid)-1 && v == len(grid[0])-1 {
		return true
	}
	for k := 0; k < 4; k++ {
		i := u + dr[k]
		j := v + dc[k]
		if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) && !visited[i][j] && mid <= grid[i][j] {
			if canDFS(grid, mid, i, j) {
				return true
			}
		}
	}
	return false
}
