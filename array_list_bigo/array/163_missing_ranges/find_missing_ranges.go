package missing_ranges

/*
	problem: https://leetcode.com/problems/missing-ranges/

	if first num != lower => add lower at the head of nums
	if last num != upper => add upper to the tail of nums

	integrate all num in nums, add [nums[i] + 1, nums[i+1]-1] to result

	time complexity: O(n). N is the number of nums
*/
