package n_ary_tree_level_order_traversal

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	input := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}
	expectation := [][]int{{1}, {3, 2, 4}, {5, 6}}
	assert.Equal(t, expectation, levelOrder(input))
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
