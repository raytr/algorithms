package run_sum_of_1_day_array

/*
problem: https://leetcode.com/problems/running-sum-of-1d-array/
prefix[i] = prefix[i-1] + A[i]
Complexity: O(n)
*/
func runningSum(nums []int) []int {

	n := len(nums)
	res := make([]int, 0, n)
	res = append(res, nums[0])

	for i := 1; i < n; i++ {
		res = append(res, res[i-1]+nums[i])
	}
	return res
}
