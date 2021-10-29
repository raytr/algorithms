package minimum_knight_moves

import "fmt"

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
	queue := [][]int{{0, 0}}
	dis := 0
	visited := make(map[string]bool)

	for len(queue) > 0 {
		for _, node := range queue {
			for k := 0; k < 8; k++ {
				i := node[0] + dr[k]
				j := node[1] + dc[k]
				if i == x && j == y {
					return dis + 1
				}
				key := fmt.Sprintf("%v,%v", i, j)
				if i >= -300 && i <= 300 && j >= -300 && j <= 300 && !visited[key] {
					visited[key] = true
					queue = append(queue, []int{i, j})
				}
			}
			queue = queue[1:]
		}
		dis++
	}
	return -1
}
