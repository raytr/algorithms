package number_of_islands

var dr = []int{0, 1, 0, -1}
var dc = []int{1, 0, -1, 0}
var count = 0
var visited = make([][]bool, 0)

/*
	problem: https://leetcode.com/problems/number-of-islands/

	to move next point, using 2 array:
			var dr = []int{0,1,0,-1}
			var dc = []int{1,0,-1,0}
	=> up: [-1, 0]
		down: [1,0]
		left: [0,-1]
		right: [0,1]

	iterate all elements in grid
	if visit[i+j] == false && grid is land (grid[i][j] == 1 => dfs
	in the end => return count

	bfs complexity: 1 note visited 1 time ; 1 edge visited at most 2 time

	complexity: O(len of grid)
*/

func numIslands(grid [][]byte) int {
	for r := 0; r < len(grid); r++ {
		visited = append(visited, []bool{})
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 1 && !visited[r][c] {
				count++
				dfs(grid, r, c)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, r, c int) {
	visited[r][c] = true
	for k := 0; k < 4; k++ {
		i := r + dr[k]
		j := c + dc[k]
		if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) && grid[i][j] == 1 && !visited[r][c] {
			dfs(grid, i, j)
		}
	}
}
