package longest_common_subsequence

/*
	problem: https://leetcode.com/problems/longest-common-subsequence/

	   state:
	       f[i][j] is the len of longest common sequence when considering s[0...i], t[0...j]

	   if s[i] == t[j] => f[i][j] == f[i-1][j-1]
	   if s[i] != t[j] =>
	       f[i][j] = max(f[i-1][j], f[i][j-1])

	time complexity : O(m*n)
	space complexity: O(m*n)
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)

	// if one in 2 text is empty => return 0
	if n*m == 0 {
		return 0
	}

	// create fx is 2D array and set all element of fx is 0
	// O(n)
	fx := make([][]int, n+1, n+1)
	for i := 0; i <= n; i++ {
		fx[i] = make([]int, m+1, m+1)
	}

	for i := 1; i <= n; i++ { //O(n)
		for j := 1; j <= m; j++ { //O(m)
			if text1[i-1] == text2[j-1] {
				fx[i][j] = fx[i-1][j-1] + 1
			} else {
				fx[i][j] = Max(fx[i-1][j], fx[i][j-1])
			}
		}
	}
	return fx[n][m]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
