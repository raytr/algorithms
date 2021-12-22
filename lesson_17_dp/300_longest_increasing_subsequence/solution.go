package longest_increasing_subsequence

/*
	problem: https://leetcode.com/problems/longest-increasing-subsequence/

	   state:
	       f[i]  ~ the lengh of lengest increase sequence that ended at i
	   relation between problem and its subproblems
	       f[i] = f[j]+1 with A[j] is sequence i > j and A[i]>A[j] and ended at i



	   f[1] = 1
	   f[2] = 1
	   f[3] = 1
	   f[4] = 2
*/

func lengthOfLIS(nums []int) int {

	n := len(nums)
	f := make(map[int]int)

	//base case
	f[0] = 1

	// state: find the len increasing sequence that ended at A[i] (1 -> n)
	for i := 1; i < n; i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				f[i] = Max(f[i], f[j]+1)
			}
		}
	}

	// find the len of the longest increasing sequence of the whole nums
	res := f[0]
	for i := 0; i < n; i++ {
		res = Max(res, f[i])
	}

	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
