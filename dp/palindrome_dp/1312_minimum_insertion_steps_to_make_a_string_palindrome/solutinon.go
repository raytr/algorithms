package minimum_insertion_steps_to_make_a_string_palindrome

/*
   the ideal is find the longest palindrome substring, after that len(s) - len(subs)
*/

func minInsertions(s string) int {
	n := len(s)
	f := make([][]bool, n, n)
	for i := 0; i < n; i++ {
		f[i] = make([]bool, n, n)
		f[i][i] = true
		if i+1 < n {
			f[i][i+1] = s[i] == s[i+1]
		}
	}

	// build f
	for leng := 3; leng <= n; leng++ {
		for left := 0; left < n-leng+1; left++ {
			right := left + leng - 1
			if s[left] == s[right] {
				f[left][right] = f[left+1][right-1]
			} else if s[left] != s[right] {
				f[left][right] = false
			}
		}
	}

	longest := 0 //1st is total character, 2nd is left index, 3rd is right index
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if f[i][j] && j-i > longest {
				longest = j - i + 1
			}
		}
	}

	return len(s) - longest
}
