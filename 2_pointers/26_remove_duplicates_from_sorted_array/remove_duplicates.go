package remove_duplicates_from_sorted_array

// problem: https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {

	/*
	   we have 2 pointers f & s at the first element
	    move f, if nums[f] > nums[s] => nums[s+1]=nums[f]
	    similar to f go over all num in nums
	*/
	if len(nums) == 0 {
		return 0
	}
	s := 0
	for f := 1; f < len(nums); f++ {
		if nums[f] > nums[s] {
			s++
			nums[s] = nums[f]
		}
	}
	return s + 1
}
