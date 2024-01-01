package majority_element

/*
	problem: https://leetcode.com/problems/majority-element
*/

func majorityElement(nums []int) int {
	maxNum, maxFreq := nums[0], 1

	freqM := make(map[int]int)

	for _, num := range nums {
		freqM[num]++
		if maxFreq < freqM[num] {
			maxNum = num
			maxFreq = freqM[num]
		}
	}

	return maxNum
}
