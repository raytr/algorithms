package house_robber

/*
	problem: https://leetcode.com/problems/house-robber/

Define state:

	f[i] ~ the max money can rob up to and including house i

Find relation
	considering house i:
		if rob house i: f[i] = f[i-2] + money[i]
		if not rob house i: f[i] = f[i-1]
	so f[i] = max(f[i-1], f[i-2] + money[i])
Base case:
	f[0] = money[0]
	f[1] = max(money[0], money[1])

*/

var f map[int]int

func rob(nums []int) int {
	n := len(nums)
	f = make(map[int]int)

	for i := 2; i <= n; i++ {
		f[i] = -1
	}

	f[0] = 0
	f[1] = nums[0]
	return call(nums, len(nums))
}

func call(nums []int, i int) int {
	if f[i-2] == -1 {
		f[i-2] = call(nums, i-2)
	}
	rob := f[i-2] + nums[i-1]

	if f[i-1] == -1 {
		f[i-1] = call(nums, i-1)
	}
	noRob := f[i-1]

	return Max(rob, noRob)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
