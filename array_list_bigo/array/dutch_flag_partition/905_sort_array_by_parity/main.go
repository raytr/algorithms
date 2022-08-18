package main

//problem: https://leetcode.com/problems/sort-array-by-parity

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity([]int{104, 80, 149, 65}))
	fmt.Println(sortArrayByParity([]int{}))
}
func sortArrayByParity(nums []int) []int {
	n := len(nums)

	boundary := 0
	for i := 0; i < n; i++ {
		if nums[i]%2 == 0 {
			nums[i], nums[boundary] = nums[boundary], nums[i]
		}
	}

	boundary = n - 1
	for i := n - 1; i >= 0; i-- {
		if nums[i]%2 == 1 {
			nums[i], nums[boundary] = nums[boundary], nums[i]
		}
	}

	return nums
}
