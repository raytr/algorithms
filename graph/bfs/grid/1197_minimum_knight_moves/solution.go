package minimum_knight_moves

/*
problem: https://leetcode.com/problems/minimum-knight-moves/
create the knightMovement:
	dr := [2, 1, -2, 1, -2, -1, 2, -1]
	dc := [1, 2, 1, -2, -1, -2, -1, 2]

use bfs
	queue [][]int
	queue.push(root)

	while !queue.IsEmpty {
		count++
		for node range of queue {
			for k:=0; k<8; k++ {
				i := dr[k]+node[0]
				j := dc[k]+node[0]
				if i==x && j == y => return count+1
				mark the node is visited
				queue.push([i,j])
				dequeue
			}
		}
	}
	return -1

complexity: O(n)
*/

var dr = []int{2, 1, -2, 1, -2, -1, 2, -1}
var dc = []int{1, 2, 1, -2, -1, -2, -1, 2}

func minKnightMoves(x int, y int) int {
	if x == 0 && y == 0 {
		return 0
	}
	queue := [][]int{{0, 0}}
	dis := make([][]int, 601)
	for i := 0; i < 601; i++ {
		dis[i] = make([]int, 601)
	}
	visited := make([][]bool, 601)
	for i := 0; i < 601; i++ {
		visited[i] = make([]bool, 601)
	}
	visited[300][300] = true
	dis[300][300] = 0
	for len(queue) > 0 {
		u := queue[0][0]
		v := queue[0][1]
		queue = queue[1:]
		if u == x && v == y {
			return dis[u+300][v+300]
		}
		for k := 0; k < 8; k++ {
			i := u + dr[k]
			j := v + dc[k]
			if i >= -300 && i <= 300 && j >= -300 && j <= 300 && !visited[i+300][j+300] {
				dis[i+300][j+300] = dis[u+300][v+300] + 1
				visited[i+300][j+300] = true
				queue = append(queue, []int{i, j})
			}
		}
	}
	return -1
}
