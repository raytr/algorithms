package max_sum_array

func maxSubArray(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	prefixSums := []int{0}
	max := -999999
	min := 0

	l := len(nums)
	for i := 0; i < l; i++ {
		prefixSums = append(prefixSums, prefixSums[i-1+1]+nums[i])
		if max < prefixSums[i+1] {
			max = prefixSums[i+1]
		}
		if min > prefixSums[i+1] {
			min = prefixSums[i+1]
		}
	}
	return max - min
}
