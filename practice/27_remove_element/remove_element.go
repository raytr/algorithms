package remove_element

/*
	problem: https://leetcode.com/problems/remove-element
	move reader and writer until meet the val. Keep moving reader until meet the elem # val
	=> swap nums[writer] & nums[reader]

	time complexity: O(n)
*/

func removeElement(nums []int, val int) int {
	n := len(nums)
	writer := 0

	for reader := 0; reader < n; reader++ {
		if nums[reader] != val {
			nums[writer] = nums[reader]
			writer++
		}
	}

	return writer
}
