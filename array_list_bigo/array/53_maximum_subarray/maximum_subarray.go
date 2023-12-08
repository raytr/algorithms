package maximum_subarray

/*
	problem: https://leetcode.com/problems/maximum-subarray/

    curSUm, curNum = firstNum
    interate each number
        if curNum > curSum =>curSum = nums[i]
        if curSUm > max => update max

    return max

    clarify: what are maximum and minimum number
    what is maximum of number in this arr

	time complexity: O(n)
	extra space: constant

*/

func maxSubArray(nums []int) int {
	curSum, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		curSum = getMax(nums[i], curSum+nums[i])
		max = getMax(max, curSum)
	}
	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
