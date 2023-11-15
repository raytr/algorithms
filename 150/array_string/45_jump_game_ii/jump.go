package jump

/*
		problem: https://leetcode.com/problems/jump-game-ii


		end pointer is created to mark the end of 1 jump step
		far pointer is the maximum index we can reach
		when more curIndex by 1, we have far = max(current far, i + nums[i])

		[2  3  1   1   4]
	     i     fe					cause i == e => update value of end pointer with value of far pointer; count++
	        i  e       f
		       i       fe => because we finish a jumb -> update e = f; count++
	               i
						ife
*/

func jump(nums []int) int {
	n, end, far, count := len(nums), 0, 0, 0

	for i := 0; i < n-1; i++ {
		far = getMax(far, i+nums[i])

		if i == end {
			end = far
			count++
		}
	}

	return count
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
