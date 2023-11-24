package k_smallest_element_in_an_array

//problem: give and array, find k smallest elements

// complexity O(len(nums))
func FindKNumsSmallest(nums []int, k int) []int {
	result := make([]int, 0, k)
	kthNum := findTheKthSmallest(nums, k)
	for _, num := range nums {
		if num <= kthNum {
			result = append(result, num)
		}
	}

	return result
}

func findTheKthSmallest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) <= k {
		return nums[len(nums)-1]
	}

	kIndex := k - 1
	pivot := nums[kIndex]

	//use floyd algo to all nums <= to left and all nums > pivot to right
	boundary := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < pivot {
			nums[i], nums[boundary] = nums[boundary], nums[i]
			boundary++
		}
	}

	boundary = len(nums) - 1
	for i := boundary; i >= 0; i-- {
		if nums[i] > pivot {
			nums[i], nums[boundary] = nums[boundary], nums[i]
			boundary--
		}
	}

	//find pivot index
	pivotIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == pivot {
			pivotIndex = i
		}
	}

	if pivotIndex == kIndex {
		return pivot
	}
	if pivotIndex > kIndex {
		left := append([]int{}, nums[:pivotIndex]...)
		return findTheKthSmallest(left, k)
	}
	if pivotIndex < kIndex {
		k = k - pivotIndex + 1
		right := append([]int{}, nums[pivotIndex+1:]...)
		return findTheKthSmallest(right, kIndex)
	}
	return 0
}
