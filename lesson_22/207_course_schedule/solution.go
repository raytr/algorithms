package course_schedule

/*

	problem: https://leetcode.com/problems/course-schedule/


    we revert this problem to grapth problem
        => if our graph has infinity loop => return false
        => otherwise => return true

    use DFS and keep track all edges by 2D array
    if meet visited that has value is true => return false

    return true
*/

var viseted [][]bool

func canFinish(numCourses int, prerequisites [][]int) bool {

	adjList := buildAdjs(numCourses, prerequisites)

	for i := 0; i < numCourses; i++ {
		//renew visited
		viseted = make([][]bool, numCourses, numCourses)
		for i := 0; i < numCourses; i++ {
			viseted[i] = make([]bool, numCourses, numCourses)
		}
		if !dfs(adjList, i) {
			return false
		}
	}
	return true
}

func dfs(adjs map[int][]int, curNode int) bool {
	for _, n := range adjs[curNode] {
		if !viseted[curNode][n] {
			viseted[curNode][n] = true
			if !dfs(adjs, n) {
				return false
			}
			viseted[curNode][n] = false //to refresh result, we just check 1 round to find cycle
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
