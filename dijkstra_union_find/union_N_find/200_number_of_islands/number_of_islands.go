package number_of_islands

/*
problem: https://leetcode.com/problems/number-of-islands/

	    first, build dr, dc
	    func dfs{
					separate to 4 direction
						if r >= 0 && r < totalR && c >= 0 && c < totalC && grid[r][c] == 1 {
						visited[r][c] = true
						dfs
	            }
	    }


	    totalR = len(grid)
	    totalC = len(grid[0])
	    for r
	        for c
	            count++
	            dfs
*/
var (
	gVisited [][]bool
	rTotal   int
	cTotal   int
	dr       = []int{1, 0, -1, 0}
	dc       = []int{0, 1, 0, -1}
)

func numIslands(grid [][]byte) int {
	gVisited = newVisited(grid)
	rTotal = len(grid)
	cTotal = len(grid[0])
	count := 0

	for r := 0; r < rTotal; r++ {
		for c := 0; c < cTotal; c++ {
			if grid[r][c] == '1' && !gVisited[r][c] {
				count++
				dfs(grid, r, c)
			}
		}
	}
	return count
}

func dfs(grid [][]byte, r, c int) {
	for k := 0; k < 4; k++ {
		u := r + dr[k]
		v := c + dc[k]
		if u >= 0 && u < rTotal && v >= 0 && v < cTotal && grid[u][v] == '1' && !gVisited[u][v] {
			gVisited[u][v] = true
			dfs(grid, u, v)
		}
	}
}

func newVisited(grid [][]byte) [][]bool {
	visited := make([][]bool, 0, len(grid))
	for i := 0; i < len(grid); i++ {
		visited = append(visited, make([]bool, len(grid[i]), len(grid[i])))
	}
	return visited
}
