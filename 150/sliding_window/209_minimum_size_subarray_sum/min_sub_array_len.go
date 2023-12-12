package minimum_size_subarray_sum

import "math"

/*
	problem: https://leetcode.com/problems/minimum-size-subarray-sum
	time complexity: because we go thought all number twise => O(N)
*/

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left := 0
	sum := 0
	minRange := math.MaxInt32

	for right := 0; right < n; right++ {
		sum += nums[right]
		for sum >= target {
			minRange = getMin(minRange, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if minRange != math.MaxInt32 {
		return minRange
	}
	return 0
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
