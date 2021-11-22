package network_delay_time

import (
	"container/heap"
	"math"
)

//func main() {
//	//fmt.Println(networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2))
//	//fmt.Println(networkDelayTime([][]int{{1, 2, 1}}, 2, 2))
//	//fmt.Println(networkDelayTime([][]int{{1, 2, 1}, {2, 1, 3}}, 2, 2))
//	//fmt.Println(networkDelayTime([][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 4}}, 3, 1))
//	//fmt.Println(networkDelayTime([][]int{{1,2,1},{2,3,7},{1,3,4},{2,1,2}}, 3, 1))
//	fmt.Println  (networkDelayTime([][]int{{1,2,1},{2,3,7},{1,3,4},{2,1,2}}, 4, 1))
//}

/*
	problem: https://leetcode.com/problems/network-delay-time/
*/

func networkDelayTime(times [][]int, n int, k int) int {
	adjM := buildAdjs(times, n)
	if adjM == nil {
		return -1
	}
	max := 0

	//init visited
	visited := make([][]bool, (n+1)*(n+1))
	for i := 1; i <= n+1; i++ {
		visited[i] = make([]bool, n+1) // 4 because we start from 1
	}

	//init queue
	priorityQ := initHeap()
	heap.Push(priorityQ, []int{k, 0})

	//init d
	d := make(map[int]int)
	for i := 1; i <= n; i++ {
		d[i] = math.MaxInt32
	}
	d[k] = 0

	for priorityQ.Len() > 0 {
		u := heap.Pop(priorityQ).([]int)[0]

		for _, pair := range adjM[u] {
			v := pair[0]
			w := pair[1]
			if !visited[u][v] {
				if d[v] > d[u]+w {
					d[v] = d[u] + w
				}
				heap.Push(priorityQ, []int{v, d[v]})
				visited[u][v] = true
			}
		}
	}
	if len(visited) < n {
		return -1
	}

	//check all of nodes was visted
	for _, t := range times {
		if !visited[t[0]][t[1]] {
			return -1
		}
	}

	for _, v := range d {
		max = Max(max, v)
	}

	return max
}

func buildAdjs(time [][]int, n int) map[int][][]int {
	adjM := make(map[int][][]int)
	nodeVisited := make(map[int]bool)
	for _, t := range time {
		nodeVisited[t[0]] = true
		nodeVisited[t[1]] = true
		if _, exist := adjM[t[0]]; !exist {
			adjM[t[0]] = make([][]int, 0)
		}
		adjM[t[0]] = append(adjM[t[0]], []int{t[1], t[2]})
	}

	if len(nodeVisited) < n {
		return nil
	}
	return adjM
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//--- heap ---

type TimeHeap [][]int

func (h TimeHeap) Len() int           { return len(h) }
func (h TimeHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h TimeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func initHeap() *TimeHeap {
	h := TimeHeap{}
	heap.Init(&h)
	return &h
}

func (h *TimeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *TimeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
