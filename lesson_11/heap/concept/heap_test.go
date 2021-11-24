package heap_concept

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := initMinHeap([]int{10, 2, 5, 7, 4, 8})
	assert.Equal(t, 2, heap.Pop(h).(int))
	heap.Push(h, 3)
	assert.Equal(t, 3, heap.Pop(h).(int))
}

func initMinHeap(nums []int) *MinHeap {
	h := MinHeap(nums)
	heap.Init(&h)
	return &h
}

func TestMaxHeap(t *testing.T) {
	h := initMaxHeap([]int{2, 5, 7, 11, 17, 4, 8})
	assert.Equal(t, 17, heap.Pop(h).(int))
	heap.Push(h, 3)
	assert.Equal(t, 11, heap.Pop(h).(int))
}

func initMaxHeap(nums []int) *MaxHeap {
	h := MaxHeap(nums)
	heap.Init(&h)
	return &h
}
