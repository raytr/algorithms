package jump_game

/*
	problem: https://leetcode.com/problems/jump-game

	             i
	[2  3  1  1  4]
	backtracking
	lastPos is the tail of nums
	i+nums[i] >= lastPos => mean: at index i can jump to lastPos
		-> keep check the previous value (dynamic programming)

	time complexity is: O(N), N is the length of nums
	space complexity is : O(1)

*/

func canJump(nums []int) bool {
	n := len(nums)
	lastPos := n - 1

	for i := n - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}

	return lastPos == 0
}
