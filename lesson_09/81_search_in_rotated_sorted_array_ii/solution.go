package search_in_rotated_sorted_array_ii

/*

problem:https://leetcode.com/problems/search-in-rotated-sorted-array-ii/

l           m           h
2   2   2   2   2   1   2

l     m        h
5  6  7  8  0  1

l        m           h
7  0  1  2  3  4  5  6

while lo <= hi
because has duplicate number in nums => if hi == lo => hi = hi-1, lo = lo+1
if mid == target => return true
if mid > lo || (mid == lo && mid < hi)) => left part doesn't has inflection
	=> if target >= lo => hi = mid-1

if mid < hi => right part doesn't have inflection
	=> if target <= hi => lo = mid+1

*/

func search(nums []int, target int) bool {
	if len(nums) == 1 && nums[0] == target {
		return true
	} else if len(nums) == 1 {
		return false
	}

	hi := len(nums) - 1
	lo := 0
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return true
		}
		if nums[hi] == nums[lo] && target != nums[lo] {
			hi = hi - 1
			lo = lo + 1
		} else if nums[mid] > nums[lo] || (nums[mid] == nums[lo] && nums[mid] > nums[hi]) {
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
	return false
}
