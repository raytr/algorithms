package minimum_size_subarray_sum

import (
	"math"
)

/*
	problem: https://leetcode.com/problems/minimum-size-subarray-sum
	time complexity: because we go thought all number twise => O(N)
 */

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans := math.MaxInt32
	left := 0
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		for sum >= target {
			ans = min(ans, i+1-left)
			sum -= nums[left]
			left++
		}
	}
	if ans != math.MaxInt32  {
		return ans
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}