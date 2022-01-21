package network

/*
   build adj list
   solve by BFS
	build visited map[int]bool
	build adj list => finded the source
	create queue and add source to queue
	with each city in queue

*/
func distanceCount(T []int32) []int32 {

	n := len(T)
	//revert to int
	t := make([]int, 0, n)
	for _, num := range T {
		t = append(t, int(num))
	}
	adjList, source := buildAdjList(t) //O(n)
	visited := buildVisited(n)         //O(n)

	disMapFromSrc := make(map[int]int)

	//create queue
	queue := []int{source}

	for len(queue) > 0 {
		// pop the first node
		cur := queue[0]
		visited[cur] = true
		queue = queue[1:]
		for _, near := range adjList[cur] {
			if !visited[near] {
				disMapFromSrc[near] = disMapFromSrc[cur] + 1
				queue = append(queue, near)
			}
		}
	}

	//build freqmap
	disFreqMap := make(map[int]int)
	for _, v := range disMapFromSrc {
		disFreqMap[v]++
	}

	// build res based on disMap
	res := make([]int32, n-1, n-1)
	for i := 1; i < n; i++ {
		res[i-1] = int32(disFreqMap[i])
	}
	return res
}

func buildVisited(n int) map[int]bool {
	visited := make(map[int]bool)
	for i := 0; i < n; i++ {
		visited[i] = false
	}
	return visited
}

func buildAdjList(T []int) (map[int][]int, int) {
	n, source := len(T), -1

	adjList := make(map[int][]int)
	for i := 0; i < n; i++ {
		adjList[i] = make([]int, 0, n)
	}

	for i, v := range T {
		if i == v {
			source = v
		} else {
			adjList[i] = append(adjList[i], v)
			adjList[v] = append(adjList[v], i)
		}
	}

	return adjList, source
}
