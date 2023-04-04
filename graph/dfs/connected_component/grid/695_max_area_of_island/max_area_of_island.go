package max_area_of_island

/*
	problem: https://leetcode.com/problems/max-area-of-island/
	time complexity: O(n)
    space complexity: O(n) (need space to store gVisited)
*/

var (
	dr      = []int{1, 0, -1, 0}
	dc      = []int{0, 1, 0, -1}
	cTotal  int
	rTotal  int
	visited [][]bool
)

func maxAreaOfIsland(grid [][]int) int {
	rTotal = len(grid)
	cTotal = len(grid[0])
	count := 0
	visited = newVisited(grid)

	for r := 0; r < rTotal; r++ {
		for c := 0; c < cTotal; c++ {
			b := grid[r][c]
			if visited[r][c] == false && b == 1 {
				count++
				visited[r][c] = true
				dfs(grid, r, c)
			}
		}
	}

	return count
}

func dfs(grid [][]int, r, c int) {
	for k := 0; k < 4; k++ {
		u := r + dr[k]
		v := c + dc[k]
		if u >= 0 && u < rTotal && v >= 0 && v < cTotal && !visited[u][v] && grid[u][v] == 1 {
			visited[u][v] = true
			dfs(grid, u, v)
		}
	}
}

func newVisited(grid [][]int) [][]bool {
	v := make([][]bool, 0, len(grid))
	for i := 0; i < len(grid[0]); i++ {
		v = append(v, make([]bool, len(grid[0]), len(grid[0])))
	}
	return v
}
