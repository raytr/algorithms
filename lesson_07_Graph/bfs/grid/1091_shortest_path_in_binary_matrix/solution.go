package shortest_path_in_binary_matrix

/*
problem: https://leetcode.com/problems/shortest-path-in-binary-matrix/
define :
dr := []int{0,1,0,-1, -1, 1, 1, -1}
dc := []int{1,0,-1,0, -1, 1, -1, 1}

queue []*Node => queue.push(root)

while !queue.IsEmpty{
	count++
	for node := range queue{
		node => for 8 time represent 8 directions{
		i = q[0][0] + dr
		j = q[0][1] + dc
		if i,j = m-1, n-1, break
		if ! out of rectangle && !visited && val == 1 => add to queue
	}
}






	return count
}

*/
var dr = []int{0, 1, 0, -1, -1, 1, 1, -1}
var dc = []int{1, 0, -1, 0, -1, 1, -1, 1}

func shortestPathBinaryMatrix(grid [][]int) int {
	maxR := len(grid) - 1
	maxC := len(grid[0]) - 1
	if grid[0][0] == 1 || grid[maxR][maxC] == 1 {
		return -1
	}
	if len(grid) == 1 && len(grid[0]) == 1 && grid[0][0] == 0 {
		return 1
	}

	q := [][]int{{0, 0}}
	dis := 0
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}

	for len(q) > 0 {
		for _, node := range q {
			for k := 0; k < 8; k++ {
				i := node[0] + dr[k]
				j := node[1] + dc[k]

				if i == maxR && j == maxC {
					//+2 because :
					//+ 1 for edge (this step just route search and distance not increase yet.
					//+ 1 because this problem return node (= edge + 1)
					return dis + 2
				}
				if i >= 0 && i <= maxR && j >= 0 && j <= maxC && !visited[i][j] && grid[i][j] == 0 {
					visited[i][j] = true
					q = append(q, []int{i, j})
				}
			}
		}
		dis++
		q = q[1:]
	}
	return -1
}
