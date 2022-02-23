package find_minimum_in_rotated_sorted_array

/*
	problem: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/


First at all, find smallest element
Func findIdxOfSmallest {

l	       m 	 	     h
3. 4,  5,  6,  7,  0  1  2
find continuous increase part

if m < m-1 && m < m+1 => mid is smallest => return mid (index)

Mid >= lo => continuous is left part => lo is smallest number in left part
	If [lo] > [hi] => all of element in left are not smallest number => lo = mid + 1
	Otherwise => hi = mid-1

L.       m           h
7. 0. 1. 2. 3. 4. 5. 6
Mid < 1o => continuous is right part is right part => mid is smallest number in right part
	If [mid] > [mid - 1] => all of element in right part doesn’t smallest number => hi = mid – 1

}

Return len - 1

*/

func findMin(nums []int) int {
	hi := len(nums) - 1
	lo := 0

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if mid == lo {
			if nums[lo] < nums[hi] {
				return nums[lo]
			} else {
				return nums[hi]
			}
		}
		if nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1] {
			return nums[mid]
		}

		if nums[mid] >= nums[lo] {
			if nums[lo] > nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else {
			if nums[mid] > nums[mid-1] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
	}
	return nums[0]
}
