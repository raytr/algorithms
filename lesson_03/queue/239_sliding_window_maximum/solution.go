package main

import "fmt"

// use the dequeue to maintain index of all element of sliding window
// iterate over array from left to right
// - remove all of indexes of smaller than the current from dequeue
//		- because they can not be biggest => don't need to care them
// keep only the indexes of element from current slide window
// A[0] in sliding window always is biggest

/*
	Large -> small
	0   1.     2.    3.    4.   5.  6.   7
	1,  3,     -1, - 3,    5,    3, 6,   7
	-3 < 3 => ok push
	 -3 check index 3 in qIndex: => yes: remove |  No : continue
	5 > -3 => remove queue, push 5 => check idex 2 in queue => yes: remove | no : continue

	complexity: O(n)
*/

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
}
func maxSlidingWindow(nums []int, k int) []int {
	var qIdx, qVal, maxElement []int
	for i := 0; i < len(nums); i++ {

		//check i-k is contained in qIdx => yes: remove | no: do nothing
		if i >= k && i-k == qIdx[0] {
			qIdx = qIdx[1:]
			qVal = qVal[1:]
		}

		//remove all element smaller than or equal nums[i]
		for len(qVal) > 0 && nums[i] > qVal[len(qVal)-1] {
			qIdx = qIdx[:len(qIdx)-1]
			qVal = qVal[:len(qVal)-1]
		}

		//add nums[i] to queue
		qIdx = append(qIdx, i)
		qVal = append(qVal, nums[i])
		if i >= k-1 {
			maxElement = append(maxElement, qVal[0])
		}
	}

	return maxElement
}
