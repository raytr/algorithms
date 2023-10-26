package course_schedule

/*

	problem: https://leetcode.com/problems/course-schedule/


    we revert this problem to grapth problem
        => if our graph has infinity loop => return false
        => otherwise => return true

	dfsIsCycle to find the node that adjList[node].len == 0 => delete this node on djs list and in visited
	if meet a node has vited[node] == true => false

	otherwise => return true

*/

//var visited []bool

//func canFinish(numCourses int, prerequisites [][]int) bool {
//
//	adjList := buildAdjs(numCourses, prerequisites)
//	visited = make([]bool, numCourses, numCourses)
//
//	for i := 0; i < numCourses; i++ {
//		if dfsIsCycle(adjList, i) {
//			return false
//		}
//	}
//	return true
//}
//
//func dfsIsCycle(adjs map[int][]int, curNode int) bool {
//	visited[curNode] = true
//	if len(adjs[curNode]) == 0 {
//		return false
//	}
//	for _, node := range adjs[curNode] {
//		if visited[node] {
//			return true
//		}
//		if !dfsIsCycle(adjs, node) {
//			delete(adjs, node)
//			visited[node] = false
//		}
//	}
//	return false
//}

var visited []bool

func canFinish(numCourses int, prerequisites [][]int) bool {

	//visited = make([]bool, numCourses, numCourses)
	adjList := buildAdjs(numCourses, prerequisites)

	for i := 0; i < numCourses; i++ {
		//renew visited
		visited = make([]bool, numCourses, numCourses)
		//for i := 0; i < numCourses; i++ {
		//	visited[i] = make([]bool, numCourses, numCourses)
		//}
		if !dfs(adjList, i) {
			return false
		}
	}
	return true
}

func dfs(adjs map[int][]int, curNode int) bool {
	visited[curNode] = true
	for _, n := range adjs[curNode] {
		if !visited[n] {
			visited[n] = true
			if !dfs(adjs, n) {
				return false
			}
			visited[n] = false //to refresh result, we just check 1 round to find cycle
		} else {
			return false
		}
	}
	return true
}
func buildAdjs(numCourses int, prerequisites [][]int) map[int][]int {
	adjs := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		adjs[i] = make([]int, 0, len(prerequisites))
	}
	for _, req := range prerequisites {
		adjs[req[0]] = append(adjs[req[0]], req[1])
	}
	return adjs
}

//func buildAdjs(numCourses int, prerequisites [][]int) map[int][]int {
//	adjs := make(map[int][]int)
//	for i := 0; i < numCourses; i++ {
//		adjs[i] = make([]int, 0, len(prerequisites))
//	}
//	for _, req := range prerequisites {
//		adjs[req[0]] = append(adjs[req[0]], req[1])
//	}
//	return adjs
//}
