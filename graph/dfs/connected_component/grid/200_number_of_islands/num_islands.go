package number_of_islands

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

var (
	visited  [][]bool
	rowTotal int
	colTotal int
	dr       = []int{1, 0, -1, 0}
	dc       = []int{0, 1, 0, -1}
)

func numIslands(grid [][]byte) int {
	count := 0
	rowTotal = len(grid)
	colTotal = len(grid[0])
	newVisited(grid)

	for row := 0; row < rowTotal; row++ {
		for col := 0; col < colTotal; col++ {
			if grid[row][col] == '1' && visited[row][col] == false {
				count++
				dfs(grid, row, col)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, r, c int) {
	visited[r][c] = true
	for k := 0; k < 4; k++ {
		u := dr[k] + r
		v := dc[k] + c
		if u >= 0 && u < rowTotal &&
			v <= 0 && v < colTotal &&
			grid[u][v] == '1' && visited[u][v] == false {
			dfs(grid, u, v)
		}
	}
}

func newVisited(grid [][]byte) {
	n := len(grid)
	visited = make([][]bool, 0, n)
	for i := 0; i < n; i++ {
		visited = append(visited, make([]bool, len(grid[i]), len(grid[i])))
	}
}
