//package find_median_from_data_stream
package main

import (
	"container/heap"
	"fmt"
)

/*
	problem: https://leetcode.com/problems/find-median-from-data-stream/

    add:
		add to heap
    find:
		while heap.len > this.length+1 : pop heap
         if len is odd number => res = pop(); h.push(res) => return res

        if len is event number =>
            => a = h.pop(); b = h.pop() => res = a + b / 2
            => heap.push(a,b)

*/
type MedianFinder struct {
	minH   *MinHeap
	maxH   *MaxHeap
	length int
}

func main() {
	mf := Constructor()
	//	["MedianFinder","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian"]
	//	input:		[[],[6],[],[10],[],[2],[],[6],[],[5],[],[0],[],[6],[],[3],[],[1],[],[0],[],[0],[]]
	// expectation: [null,6,null,8,null,6,null,6,null,6,null,5.5,null,6.,null,5.5,null,5,null,4,null,3]
	mf.AddNum(6)
	fmt.Println(mf.FindMedian())
	mf.AddNum(10)
	fmt.Println(mf.FindMedian())
	mf.AddNum(2)
	fmt.Println(mf.FindMedian())
	mf.AddNum(6)
	fmt.Println(mf.FindMedian())
	mf.AddNum(5)
	fmt.Println(mf.FindMedian())
	mf.AddNum(0)
	fmt.Println(mf.FindMedian())
	mf.AddNum(6)
	fmt.Println(mf.FindMedian())
	mf.AddNum(3)
	fmt.Println(mf.FindMedian())
	mf.AddNum(1)
	fmt.Println(mf.FindMedian())
	mf.AddNum(0)
	fmt.Println(mf.FindMedian())
	mf.AddNum(0)
	fmt.Println(mf.FindMedian())

}

func Constructor() MedianFinder {
	minH, maxH := initHeap()
	return MedianFinder{
		minH:   &minH,
		maxH:   &maxH,
		length: 0,
	}
}

func initHeap() (MinHeap, MaxHeap) {
	minH := MinHeap{}
	heap.Init(&minH)
	maxH := MaxHeap{}
	heap.Init(&maxH)
	return minH, maxH
}

func (this *MedianFinder) AddNum(num int) {
	if this.minH.Len() == 0 && this.maxH.Len() == 0 {
		heap.Push(this.maxH, num)
	} else {
		firstMaxHeap := heap.Pop(this.maxH).(int)
		if (num) > firstMaxHeap {
			heap.Push(this.minH, num)
			heap.Push(this.maxH, firstMaxHeap)
		} else {
			heap.Push(this.minH, firstMaxHeap)
			heap.Push(this.maxH, num)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	var res float64
	if this.maxH.Len() == this.minH.Len() {
		lo := heap.Pop(this.maxH).(float64)
		hi := heap.Pop(this.minH).(float64)
		res = (hi + lo) / 2
	} else if this.maxH.Len() > this.minH.Len() {
		if this.maxH.Len() > 0 {
			r := heap.Pop(this.maxH)
			res = r.(float64)
			heap.Push(this.maxH, res)
		}
	} else {
		res = heap.Pop(this.minH).(float64)
		heap.Push(this.minH, res)
	}

	return res
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

//------------------ min heap----------------
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//------------------ min heap----------------
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n+1]
	*h = old[0 : n+1]
	return x
}
