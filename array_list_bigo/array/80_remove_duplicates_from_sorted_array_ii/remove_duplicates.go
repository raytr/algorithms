package remove_duplicates_from_sorted_array_ii

/*
	problem: https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii
*/

func removeDuplicates(nums []int) int {
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}

		if count > 2 {
			nums = removeElemAtIndex(nums, i)
			count--
			i--
		}
	}

	return len(nums)
}

func removeElemAtIndex(nums []int, removedIndex int) []int {
	nums = append(nums[:removedIndex], nums[removedIndex+1:]...)
	return nums
}
