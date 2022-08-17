package run_sum_of_1day_array

//problem: https://leetcode.com/problems/running-sum-of-1d-array/
//prefix[i] = prefix[i-1] + A[i]
//Complexity: O(n)
func runningSum(nums []int) []int {
	output := []int{nums[0]}
	n := len(nums)

	for i := 1; i < n; i++ {
		output = append(output, output[i-1]+nums[i])
	}
	return output
}
