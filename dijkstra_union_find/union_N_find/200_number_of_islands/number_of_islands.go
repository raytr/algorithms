package number_of_islands

/*

problem: https://leetcode.com/problems/number-of-islands/

    first, build dr, dc
    func dfs{
				sperate to 4 direction
					if r >=0 && r < totalR && c >=0 && c < totalC && grid[r][c] == 1 {
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
var visited [][]bool
var totalR int
var totalC int
var dr = []int{1, 0, -1, 0}
var dc = []int{0, 1, 0, -1}

func numIslands(grid [][]byte) int {
	totalR = len(grid)
	totalC = len(grid[0])
	count := 0

	//create visited
	visited = make([][]bool, totalR, totalR)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, totalC, totalC)
	}

	for r := 0; r < totalR; r++ {
		for c := 0; c < totalC; c++ {
			if grid[r][c] == byte('1') && !visited[r][c] {
				visited[r][c] = true
				count++
				dfs(grid, r, c)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, r, c int) {
	for i := 0; i < 4; i++ {
		u := r + dr[i]
		v := c + dc[i]
		if u >= 0 && u < totalR && v >= 0 && v < totalC && !visited[u][v] && grid[u][v] == byte('1') {
			visited[u][v] = true
			dfs(grid, u, v)
		}
	}
}
