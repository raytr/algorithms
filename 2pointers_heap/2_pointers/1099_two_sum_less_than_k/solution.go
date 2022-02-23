package two_sum_less_than_k

import "sort"

/*
	problem: https://leetcode.com/problems/two-sum-less-than-k/

   we have 2 pointers is l = 0 & r = len-1
   while l <= r
       if sum < k => res = sum, l++
       ortherwise is r--

   return res
*/

func twoSumLessThanK(nums []int, k int) int {

	if len(nums) < 2 {
		return -1
	}
	sort.Ints(nums)
	max, l, r := -1, 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		if sum < k {
			max = Max(max, sum)
			l++
		} else {
			r--
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
