package run_sum_of_1day_array

//problem: https://leetcode.com/problems/running-sum-of-1d-array/

//Complexity: O(n)
func runningSum(nums []int) []int {
	// first about edge cache
	if len(nums) == 1 {
		return []int{nums[0]}
	}

	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}

	return nums
}
