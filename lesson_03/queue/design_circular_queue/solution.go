package design_circular_queue

//problem: https://leetcode.com/problems/design-circular-queue/

type MyCircularQueue struct {
	head, size int
	queue      []int
}

func Constructor(k int) MyCircularQueue {
	myCircularQueue := MyCircularQueue{
		size:  0,
		head:  0,
		queue: make([]int, k, k),
	}
	return myCircularQueue
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.size++
	tail := this.GetTail()
	this.queue[tail] = value
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.size--
	this.head = (this.head + 1) % len(this.queue)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	tail := this.GetTail()
	return this.queue[tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if len(this.queue) == this.size {
		return true
	}
	return false
}

func (this *MyCircularQueue) GetTail() int {
	if this.IsEmpty() {
		return this.head
	}
	return (this.head + this.size - 1) % len(this.queue)
}
