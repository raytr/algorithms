package kth_largest_element_in_an_array

// the worst case: O(n^2)
// the best case: O(n)
func findKthLargest(nums []int, k int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	// apply Dutch flag partition
	// pivot should be the len - k
	kIndex := len(nums) - k
	pivot := nums[kIndex]

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

	pivotIndex := 0
	for i, v := range nums {
		if v == pivot {
			pivotIndex = i
		}
	}
	//if k == pivot => pivot
	// if length of nums - pivot index  > k => recurs left
	// if length of nums - pivot index  < k => recurs right
	if pivotIndex == kIndex {
		return pivot
	}
	if pivotIndex > kIndex {
		arr := append(nums[:pivotIndex])
		k = len(nums) - pivotIndex
		return findKthLargest(arr, k)
	}
	if pivotIndex < kIndex {
		arr := append([]int{}, nums[pivotIndex:]...)
		return findKthLargest(arr, k)
	}
	return 0
}
