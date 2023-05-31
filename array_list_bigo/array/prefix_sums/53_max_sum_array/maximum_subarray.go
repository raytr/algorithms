package _3_max_sum_array

import "math"

//problem: https://leetcode.com/problems/maximum-subarray/

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	max := -math.MaxInt
	min := math.MaxInt
	prefixSums := []int{0}

	for i := 1; i < len(nums); i++ {
		prefixSumNum := prefixSums[i-1] + nums[i]
		prefixSums = append(prefixSums, prefixSumNum)

		if max < prefixSumNum {
			max = prefixSumNum
		}

		if min > prefixSumNum {
			min = prefixSumNum
		}
	}

	return max - min
}
