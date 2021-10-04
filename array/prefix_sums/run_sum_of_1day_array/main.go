package run_sum_of_1day_array

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
