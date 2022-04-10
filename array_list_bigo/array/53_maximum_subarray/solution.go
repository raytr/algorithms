package maximum_subarray

/*
	problem: https://leetcode.com/problems/maximum-subarray/

   iterate each number,
       if currentSum < currentNum -> currentSum = currentNum
       max = getMax(currentSum, max)

   O(n)

*/
func maxSubArray(nums []int) int {
	max := nums[0]
	curSum := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		curSum = getMax(num, curSum+num)
		max = getMax(curSum, max)
	}
	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
