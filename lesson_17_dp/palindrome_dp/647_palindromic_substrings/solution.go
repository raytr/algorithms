package palindromic_substrings

/*
		problem: https://leetcode.com/problems/palindromic-substrings/

	   f[i][j] is boolean
	   if f[i][j] == true if s[i...j] is palindrome
	   if f[i][j] == false if s[i...j] isn't palindrome

	   if s[i] == s[j] => f[i][j] == f[i+1][j-1]
	   if s[i] == s[j] => false

	   base case
	       s[i][i] == true
	       s[i][i+1] = s[i] == s[i+1]

		time complexity O(n^2)
		space complecity O(n^2)

*/

func countSubstrings(s string) int {
	n := len(s)
	fx := make([][]bool, n, n)
	for i := 0; i < n; i++ {
		fx[i] = make([]bool, n, n)
		fx[i][i] = true
		if i+1 < n {
			fx[i][i+1] = s[i] == s[i+1]
		}
	}

	//build the fx
	for leng := 3; leng <= n; leng++ {
		for left := 0; left < n-leng+1; left++ {
			right := left + leng - 1
			if s[left] == s[right] {
				fx[left][right] = fx[left+1][right-1]
			} else if s[left] != s[right] {
				fx[left][right] = false
			}
		}
	}

	// push all fx has value == true to array
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if fx[i][j] {
				res++
			}
		}
	}

	return res
}
