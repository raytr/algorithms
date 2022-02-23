package climbing_stairs

/*
	problem: https://leetcode.com/problems/climbing-stairs/

	   to climb to the top has n steps, step[n] = step[n-1] + 1

	   step[1] = 1
	   step[2] = 2
	   step[3] = step[1] + step[2] = 3


	   state:
	       count[n] ~ number to distinc way to climb to ith step
	   relation between problem and subproblems:
	       count[i] = count[i-1] + count[i-2]
	   base case:
	       count[1] = 1
	       count[2] = 2


	runtime : O(n)
	space complexity: O(1)
*/

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	res, a, b := 0, 1, 2

	for i := 3; i <= n; i++ {
		res = a + b
		a = b
		b = res
	}
	return res
}
