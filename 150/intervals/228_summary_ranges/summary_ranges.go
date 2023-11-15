package summary_ranges

import "fmt"

/*
	problem: https://leetcode.com/problems/summary-ranges

	[0,1,2,4,5,7]

	for i <= len(nums)
		firstNum := num[i] // 0
		for nums[i+1] == nums[i] + 1
			i++ // 1

		firstNum -> num[i+1] 2 "0 -> 2"
		firstNum i+2

	time complexity: O(N), N is length of nums
	space complexity: O(N), needing store the result. The worst case is that in nums there aren't any adjacent numbers
*/

func summaryRanges(nums []int) []string {
	n := len(nums)
	result := make([]string, 0, n)

	for i := 0; i < n; {
		firstIndex := i
		if i == n-1 {
			result = append(result, fmt.Sprintf("%v", nums[i]))
			return result
		}

		for i < n-1 && nums[i]+1 == nums[i+1] {
			i++
		}

		if firstIndex == i {
			result = append(result, fmt.Sprintf("%v", nums[i]))
		} else {
			result = append(result, fmt.Sprintf("%v->%v", nums[firstIndex], nums[i]))
		}
		i += 1
	}

	return result
}
