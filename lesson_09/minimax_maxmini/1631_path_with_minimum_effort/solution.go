package path_with_minimum_effort

import "math"

/*
	problem: https://leetcode.com/problems/path-with-minimum-effort/

   increase diff => satisfied cases will decrease
   decrease diff => smaller diff sure satified => hope to bigger diff

   lo, hi = 1, 10^6

   while lo <= hi
	   mid = lo + (hi - lo)/2
	   minDiff
	   if dfs(mid) {
		   res = mid
		   lo = mid - 1
	   } else {
		   hi = mid - 1
	   }


complexity: O(n*logn)
*/

var visited [][]bool
var dr = []int{1, 0, -1, 0}
var dc = []int{0, 1, 0, -1}

func minimumEffortPath(heights [][]int) int {
	res := 0
	lo, hi := 0, int(10e6)

	for lo <= hi {
		mid := lo + (hi-lo)/2
		visited = make([][]bool, len(heights))
		for i := 0; i < len(heights); i++ {
			visited[i] = make([]bool, len(heights[i]))
		}
		visited[0][0] = true

		if dfs(heights, mid, 0, 0) {
			res = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return res
}

func dfs(heights [][]int, mid, u, v int) bool {
	visited[u][v] = true

	if u == len(heights)-1 && v == len(heights[u])-1 {
		return true
	}

	for k := 0; k < 4; k++ {
		i := u + dr[k]
		j := v + dc[k]

		if i >= 0 && j >= 0 && i < len(heights) && j < len(heights[i]) && !visited[i][j] {
			diff := int(math.Abs(float64(heights[i][j] - heights[u][v])))
			if mid >= diff {
				if dfs(heights, mid, i, j) {
					return true
				}
			}
		}
	}
	return false
}
