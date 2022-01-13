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

//var viseted []bool

//func canFinish(numCourses int, prerequisites [][]int) bool {
//
//	adjList := buildAdjs(numCourses, prerequisites)
//	viseted = make([]bool, numCourses, numCourses)
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
//	viseted[curNode] = true
//	if len(adjs[curNode]) == 0 {
//		return false
//	}
//	for _, node := range adjs[curNode] {
//		if viseted[node] {
//			return true
//		}
//		if !dfsIsCycle(adjs, node) {
//			delete(adjs, node)
//			viseted[node] = false
//		}
//	}
//	return false
//}

var viseted []bool

func canFinish(numCourses int, prerequisites [][]int) bool {

	//viseted = make([]bool, numCourses, numCourses)
	adjList := buildAdjs(numCourses, prerequisites)

	for i := 0; i < numCourses; i++ {
		//renew visited
		viseted = make([]bool, numCourses, numCourses)
		//for i := 0; i < numCourses; i++ {
		//	viseted[i] = make([]bool, numCourses, numCourses)
		//}
		if !dfs(adjList, i) {
			return false
		}
	}
	return true
}

func dfs(adjs map[int][]int, curNode int) bool {
	viseted[curNode] = true
	for _, n := range adjs[curNode] {
		if !viseted[n] {
			viseted[n] = true
			if !dfs(adjs, n) {
				return false
			}
			viseted[n] = false //to refresh result, we just check 1 round to find cycle
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
