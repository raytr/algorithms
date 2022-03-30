package find_first_and_last_position_of_element_in_sorted_array

/*
	problem : https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

   use binary seach to find the taget => add index to res
   when have a target index => check from it spread to add to res

	runtime: O(nlogN)
	space complexity: O(1)
*/

func searchRange(nums []int, target int) []int {

	//edge cases
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}

	n := len(nums)
	res := make([]int, 0, n)

	//bin search to find the target
	lo, hi := 0, n-1
	targetIndex := -1
	if nums[lo] == nums[hi] {
		if nums[lo] == target {
			return []int{0, n - 1}
		} else {
			return []int{-1, -1}
		}
	}

	for hi >= lo {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			targetIndex = mid
			res = []int{targetIndex, targetIndex}
			break
		}
		if target < nums[mid] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	if targetIndex == -1 {
		return []int{-1, -1}
	}

	//check spread
	for i := targetIndex + 1; i < n; i++ {
		if nums[i] == target {
			res[1] = i
		}
	}
	for i := targetIndex - 1; i >= 0; i-- {
		if nums[i] == target {
			res[0] = i
		}
	}
	return res
}
