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
		start := query[0]
		if !dict[start] {
			res = append(res, -1)
			continue
		}
		dest := query[1]
		if !dict[dest] {
			res = append(res, -1)
			continue
		}

		visited := make(map[string]bool)
		visited[start] = true
		res = append(res, dfs(adjList, visited, start, dest, 1))
	}

	return res
}

func dfs(adjList map[string][]location, visited map[string]bool, start, end string, curWeight float64) float64 {
	if start == end {
		return curWeight
	}

	for _, next := range adjList[start] {
		if !visited[next.dest] {
			visited[next.dest] = true
			result := dfs(adjList, visited, next.dest, end, curWeight*next.weight)
			if result != -1.0 {
				return result
			}
		}
	}

	delete(visited, start)

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
