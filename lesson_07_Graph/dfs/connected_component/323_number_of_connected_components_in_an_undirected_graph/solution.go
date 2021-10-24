package number_of_connected_components_in_an_undirected_graph

/*
	problem: https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/
	convert to adjacency list
	0 => [1]
	1 => [2, 0]
	2 => [1]
	3 => [4]
	4 => [3]


	loop from 0 -> n
	if visited[i] = false => count++
	with i == 0 => visited[0] == false => visited[0] == true
					=> map[0] = 1 => map[1] = 0, 2 => 2 not visited yet
					=> map [2] = 1(visited) => do nothing
	....
	with i == 4 => visited[4] == true => do nothing

	return count
	complexity: O(n)
*/

func countComponents(n int, edges [][]int) int {
	adjList := make(map[int][]int)

	for i := 0; i < n; i++ {
		adjList[i] = make([]int, 0, n)
	}

	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	visited := make(map[int]bool)
	count := 0

	for i := 0; i < n; i++ {
		if visited[i] == false {
			count++
		}
		dfs(adjList, i, count, visited)
	}
	return count
}

func dfs(adjList map[int][]int, curNode, count int, visited map[int]bool) {
	for _, n := range adjList[curNode] {
		if !visited[n] {
			visited[n] = true
			dfs(adjList, n, count, visited)
		}
	}
}
