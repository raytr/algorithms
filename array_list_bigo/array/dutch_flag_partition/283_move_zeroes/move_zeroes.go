package move_zeroes

func moveZeroes(nums []int) []int {
	boundary := 0
	pivot := 0

	for i, num := range nums {
		if num > pivot {
			nums[i], nums[boundary] = nums[boundary], nums[i]
			boundary++
		}
	}

	return nums
}
