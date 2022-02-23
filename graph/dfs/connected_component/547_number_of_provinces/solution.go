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

func findCircleNum(isConnected [][]int) int {
	cityMap := make(map[int][]int)
	//convert to adjList
	for i := 0; i < len(isConnected); i++ {
		cityMap[i] = make([]int, 0, len(isConnected[i]))
		for j := 0; j < len(isConnected[i]); j++ {
			if i != j && isConnected[i][j] == 1 {
				cityMap[i] = append(cityMap[i], j)
			}
			cityMap[i] = append(cityMap[i])
		}
	}

	count := 0
	visited := make(map[int]bool)
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			count++
		}
		dfs(cityMap, i, visited)
	}
	return count
}

func dfs(adjList map[int][]int, city int, visited map[int]bool) {
	for _, nearCity := range adjList[city] {
		if !visited[nearCity] {
			visited[nearCity] = true
			dfs(adjList, nearCity, visited)
		}
	}
}
