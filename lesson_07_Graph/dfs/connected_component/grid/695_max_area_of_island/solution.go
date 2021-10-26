package max_area_of_island

var dr = []int{0, 1, 0, -1}
var dc = []int{1, 0, -1, 0}

/*
	problem: https://leetcode.com/problems/max-area-of-island/



	complexity:
*/

func maxAreaOfIsland(grid [][]int) int {
	var visited = make([][]bool, 0)
	var count = 0
	for i := 0; i < len(grid); i++ {
		visited = append(visited, make([]bool, len(grid[i])))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				count++
				dfs(grid, i, j, visited)
			}
		}
	}
	return count
}

func dfs(grid [][]int, r, c int, visited [][]bool) {
	visited[r][c] = true
	for k := 0; k < 4; k++ {
		i := dr[k] + r
		j := dc[k] + c
		if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) && grid[i][j] == 1 && !visited[i][j] {
			dfs(grid, i, j, visited)
		}
	}
}
