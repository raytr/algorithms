package number_of_provinces

/*
	problem: https://leetcode.com/problems/number-of-provinces/

	build adjList:
		- cityMap key is index of isConnect, value is array findCircleNum[key][findCircleNum[k][] (remove element at i index)]
		- for len of isConnected
		- get the map
			cityMap[0]: [1]
			cityMap[1]: [0]
			cityMap[2]: []

	for all of indexs
		if visited == false => count++ => dfs
						true => dfs

	complexity: O(n)

*/

var (
	adjList map[int][]int
	visited []bool
)

func findCircleNum(isConnected [][]int) int {
	adjList = newAdjList(isConnected)
	n := len(isConnected)
	count := 0
	visited = make([]bool, n, n)

	for i := 0; i < n; i++ {
		if visited[i] == false {
			visited[i] = true
			count++
			dfs(i)
		}
	}

	return count
}

func dfs(curNode int) {
	for _, nextNode := range adjList[curNode] {
		if visited[nextNode] == false {
			visited[nextNode] = true
			dfs(nextNode)
		}
	}
}

func newAdjList(isConnected [][]int) map[int][]int {
	n := len(isConnected)
	adjList = make(map[int][]int)

	for i := 0; i < n; i++ {
		adjList[i] = make([]int, 0, len(isConnected[i]))
		for j := 0; j < len(isConnected[i]); j++ {
			if i != j && isConnected[i][j] == 1 {
				adjList[i] = append(adjList[i], j)
			}
		}
	}

	return adjList
}
