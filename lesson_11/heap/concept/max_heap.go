package heap_concept

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	root := (*h)[0]

	n := len(*h)
	lastElement := (*h)[n-1]
	(*h)[0] = lastElement
	*h = (*h)[:n-1]
	//heapify(heap, 0)

	return root
}
