package sum_of_distances_in_tree

import "fmt"

/*
https://leetcode.com/problems/sum-of-distances-in-tree/

adjention list - ex: with edge = [0,1]

	ajdList[0] = [1]
	ajdList[1] = [0]

	use bfs to solve this problem
	map has struct map[chosen node][dest]weight

	for each: sum += for

time complexity: O(n)
space complexity: O(2n)
*/
var table map[int]map[int]int
var gQueue []int

func sumOfDistancesInTree(n int, edges [][]int) []int {
	table = make(map[int]map[int]int)
	adjList := newAdjList(edges)
	result := make([]int, 0)

	for i := 0; i < n-1; i++ {
		table[i] = make(map[int]int)
		for _, nextNode := range adjList[i] {
			dfs(adjList, i, nextNode)
		}
	}

	//b := visited
	//fmt.Println(b)
	//
	//for i := 0; i < n; i++ {
	//	sum := 0
	//	for j := 0; j < n; j++ {
	//		if i != j {
	//			a := visited
	//			b := a[fmt.Sprintf("%v-%v", i, j)]
	//			sum += b
	//		}
	//	}
	//	result = append(result, sum)
	//}

	return result
}

//func tableForNode(adjList map[int][]int, source int) [][]int {
//	table[source] = make(map[int]int)
//
//	for _, nextNode := range adjList[node] {
//		if parent != nextNode {
//			dfs(adjList, node, nextNode)
//		}
//	}
//}

func dfs(adjList map[int][]int, parent, node int) {
	t := table
	fmt.Println(t)
	if len(adjList[node]) == 1 {
		table[parent][node] = 1
		return
	}
	for _, nextNode := range adjList[node] {
		if parent != nextNode {
			dfs(adjList, node, nextNode)
			table[parent][node] = table[node][nextNode] + 1
		}
	}
}

func newAdjList(edges [][]int) map[int][]int {
	adjList := make(map[int][]int)
	for _, edge := range edges {
		if _, exist := adjList[edge[0]]; !exist {
			adjList[edge[0]] = make([]int, 0)
		}
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])

		if _, exist := adjList[edge[1]]; !exist {
			adjList[edge[1]] = make([]int, 0)
		}
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	return adjList
}
