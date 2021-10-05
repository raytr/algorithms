package main

//problem: https://leetcode.com/problems/sort-array-by-parity

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}
func sortArrayByParity(nums []int) []int {
	boundary := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[boundary], nums[i] = nums[i], nums[boundary]
			boundary++
		}
	}

	boundary = len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i]%2 != 0 {
			nums[boundary], nums[i] = nums[i], nums[boundary]
			boundary--
		}
	}

	return nums
}
