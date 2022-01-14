package two_sum_ii_input_array_is_sorted

// problem: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

func twoSum(numbers []int, target int) []int {

	/*
	   2 pointers l (index 0) & r(last index),
	   while l <= r
	       if sum (l,r) == target => return l, r
	       if sum (l,r) > target descrease r
	       otherwise increase l
	*/

	l, r := 0, len(numbers)-1
	for l <= r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}
		if sum > target {
			r--
		} else {
			l++
		}
	}
	return []int{l + 1, r + 1}
}
