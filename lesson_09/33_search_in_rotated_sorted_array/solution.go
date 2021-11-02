package search_in_rotated_sorted_array

/*
problem: https://leetcode.com/problems/search-in-rotated-sorted-array

Just check in Countinuous increase part => FIND CONTINUOUS PART FIRST

l.   t    	.   m	.        h
7    8   0  1   2   3  4  5  6
M < lo => Continuous increase part is right part
	M < t < h => t in right part, otherwise: t in left part

L.      t	    m	     t       h
3   4   5   6   7    8   0   1   2
M > lo => continuous increase part is left part
	L < t < M => t in left part, otherwise: t in right part


*/

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	lo := 0
	hi := len(nums) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[lo] {
			if target >= nums[lo] && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
