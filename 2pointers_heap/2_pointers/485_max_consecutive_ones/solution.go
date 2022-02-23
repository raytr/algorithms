package max_consecutive_ones

// problem: https://leetcode.com/problems/max-consecutive-ones

func findMaxConsecutiveOnes(nums []int) int {
	/*
	   we have 2 pointers, i & j
	   i run from 0 to end the nums
	   if nums[i] == 1 => j==i
	   run j farest as possiple => j-1 < len[nums] && nums[j+1] == 1
	   max = Max(max, j-i+1)

	   complexity is O(n)
	*/
	n := len(nums)
	max, i := 0, 0
	for i < n {
		if nums[i] == 1 {
			j := i
			for j+1 < n && nums[j+1] == 1 {
				j++
			}
			if i != j {
				max = Max(max, j-i+1)
				i = j
			} else {
				max = Max(max, 1)
			}
		}
		i++
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
