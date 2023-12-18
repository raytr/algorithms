package rotate_array

/*
	problem: https://leetcode.com/problems/rotate-array
*/

func rotate(nums []int, k int) {
	count := 0
	n := len(nums)

	// do this step to make sure we rotate only 1 cycle
	k = k % n

	for count < k {
		nums[count], nums[n-k+count] = nums[n-k+count], nums[count]
		count++
	}

	nums = append(nums[:k], append(nums[n-k:n], nums[k:n-k]...)...)
}
