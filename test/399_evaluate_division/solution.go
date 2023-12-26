package evaluate_division

/*
	problem : https://leetcode.com/problems/evaluate-division

	adj

*/

type location struct {
	dest   string
	weight float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	adjList, dict := newAdjList(equations, values)
	res := make([]float64, 0)

	for _, query := range queries {
		src := query[0]
		if !dict[src] {
			res = append(res, -1)
			continue
		}
		dest := query[1]
		if !dict[dest] {
			res = append(res, -1)
			continue
		}

		visited := make(map[string]bool)
		visited[src] = true
		res = append(res, dfs(adjList, visited, src, dest, 120))
	}

	return res
}

func dfs(adjList map[string][]location, visited map[string]bool, curNode, dest string, curWeight float64) float64 {
	if curNode == dest {
		return curWeight
	}

	for _, next := range adjList[curNode] {
		if !visited[next.dest] {
			visited[next.dest] = true
			return dfs(adjList, visited, next.dest, dest, curWeight*next.weight)
		}
	}

	return -1
}

func newAdjList(equations [][]string, values []float64) (map[string][]location, map[string]bool) {
	adjList := make(map[string][]location)
	dict := make(map[string]bool)
	for i, pair := range equations {

		// add to dict
		dict[pair[0]] = true
		dict[pair[1]] = true

		if adjList[pair[0]] == nil {
			adjList[pair[0]] = make([]location, 0)
		}
		adjList[pair[0]] = append(adjList[pair[0]], location{
			dest:   pair[1],
			weight: values[i],
		})

		if adjList[pair[1]] == nil {
			adjList[pair[1]] = make([]location, 0)
		}
		adjList[pair[1]] = append(adjList[pair[1]], location{
			dest:   pair[0],
			weight: 1 / values[i],
		})
	}

	return adjList, dict
}
