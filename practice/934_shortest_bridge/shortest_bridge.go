package shortest_bridge

/*
   dx = 1 ,0, -1, 0
   dy = 0 ,1, 0, -1

   for to find 1

   queue.push(s)
   while queue.len > 0 {
       node = pop queue
       for k =0 ; k<4; k++ {

           i = grid[node[0]+dx[k]]
           j = grid[node[1]+dy[k]]
           if grid[i][j] == 1 => {
               return dis+1
           }
           if i >= 0 && i <= maxR && j >=0 && j <= maxC && !visited[i][j] && grid[i][j] == 0 {
               queue.push*(i,j)
           }

       }
       dis++
   }

   time complexity: O(n); n is all node
   space complixity: O(n); n is all node
*/

var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func shortestBridge(grid [][]int) int {
	n := len(grid)
	dis := 0
	sourceR, sourceC := -1, -1

	visited := make([][]bool, len(grid), len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]), len(grid[0]))
	}

	//find source
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				sourceR, sourceC = i, j
				break
			}
		}
		if sourceR != -1 {
			break
		}
	}

	queue := make([][]int, 0, len(grid)*len(grid))
	queue = append(queue, []int{sourceR, sourceC})

	for len(queue) > 0 {
		cur := queue[0]
		visited[cur[0]][cur[1]] = true
		queue = queue[1:]
		for k := 0; k < 4; k++ {
			i := cur[0] + dx[k]
			j := cur[1] + dy[k]
			if i >= 0 && i < n && j >= 0 && j < n && grid[i][j] == 1 && !visited[i][j] {
				//if dis == 0 {
				//	for
				//} else {
				return dis
				//}
			}
			if i >= 0 && i < n && j >= 0 && j < n && !visited[i][j] {
				queue = append(queue, []int{i, j})
			}
		}
		dis++
	}
	return 0
}
