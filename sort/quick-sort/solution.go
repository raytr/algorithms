package main

import "fmt"

func main() {
	fmt.Println(QuickSort([]int{3, 5, 8, 1, 2, 9, 4, 7, 6}))
}

func QuickSort(nums []int) []int {
	l, r := 0, len(nums)-2
	nums = qs(nums, l, r)
	return nums
}

func qs(nums []int, l, r int) []int {
	//the pivot is last of tha array
	pivot := len(nums) - 1

	//l marker will move from left to right, that will stop when reached the number greater than or equal the pivot
	for i := l; i <= len(nums)-2; i++ {
		if nums[i] > nums[pivot] {
			l = i
			break
		}
	}

	// r marker will move to left.
	// It will stop when reached the number smaller than the pivot
	for i := r; i >= 0; i-- {
		if nums[i] < nums[pivot] || i == l {
			r = i
			break
		}
	}

	nums[l], nums[r] = nums[r], nums[l]
	if l < r {
		nums = qs(nums, l, r)
	} else if l == r {
		if nums[l] > nums[pivot] {
			nums[l], nums[pivot] = nums[pivot], nums[l]
			pivot = l
		}
		leftNums := nums[:pivot+1]
		if len(leftNums) == 1 {
			leftNums = qs(leftNums, 0, len(leftNums)-2)
		}

		rightNums := nums[pivot+2:]
		rightNums = qs(rightNums, 0, len(rightNums)-2)

		nums = append(leftNums, rightNums...)
	}
	return nums
}
