package max_consecutive_ones

/*
		problem: https://leetcode.com/problems/max-consecutive-ones

	   we have 2 pointers, i & j
	   i run from 0 to end the nums
	   if nums[i] == 1 => j==i
	   run j as far as possible => j-1 < len[nums] && nums[j+1] == 1
	   max = Max(max, j-i+1)

	   complexity is O(n)
*/

func findMaxConsecutiveOnes(nums []int) int {
	n := len(nums)
	max, i := 0, 0
	for i < n {
		if nums[i] == 1 {
			j := i
			for j+1 < n && nums[j+1] == 1 {
				j++
			}
			if i != j {
				max = getMax(max, j-i+1)
				i = j
			} else {
				max = getMax(max, 1)
			}
		}
		i++
	}
	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
