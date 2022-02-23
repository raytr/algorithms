package longest_palindromic_substring

/*
	problem: https://leetcode.com/problems/longest-palindromic-substring/

   stage: f[i][j]: boolean
           f[i][j] == true if s[i...j] is palindrome
           f[i][j] == false if s[i...j] is not palindrome

   relatition: if s[i] == s[j] => f[i][j] = f[i+1][j-1]
               if s[i] != s[j] => return false

   base case:
       f[i][i] = true
       f[i][i+1] = s[i] == s[i+1]

	j = i + len - 1

	0  1  2  3
	b  b  a  b
	i     j
	   i     j
	i        j


	i = 0; l = 2; j = 2
	i = 1, l = 2; j = 3
*/

func longestPalindrome(s string) string {
	n := len(s)
	fx := make([][]bool, n, n)
	for i := 0; i < n; i++ {
		fx[i] = make([]bool, n, n)
		fx[i][i] = true
		if i < n-1 {
			fx[i][i+1] = s[i] == s[i+1]
		}
	}

	//build the fx
	for leng := 3; leng <= n; leng++ {
		for left := 0; left < n-leng+1; left++ {
			right := left + leng - 1
			if s[left] == s[right] {
				fx[left][right] = fx[left+1][right-1]
			} else {
				fx[left][right] = false
			}
		}
	}

	// find out the longest palindrome string
	longest := 0
	longestIdxs := make([]int, 2, 2)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if fx[i][j] {
				if j-i > longest {
					longest = j - i
					longestIdxs[0], longestIdxs[1] = i, j
				}
			}
		}
	}

	return s[longestIdxs[0] : longestIdxs[1]+1]
}
