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
