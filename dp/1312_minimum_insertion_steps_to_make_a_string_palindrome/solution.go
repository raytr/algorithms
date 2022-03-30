package minimum_insertion_steps_to_make_a_string_palindrome

/*
	problem: https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/

	fx[i][j]: min step to make s[i..j] palindrome
	s[i] == s[j]:
			fx[i][j] = fx[i+1][j-1]
	s[i] != s[j]
		fx[i][j] = min(fx[i+1][j] +1,
				fx[i][j-1] + 1)

	base case
	fx[i][i] = 0
	fx[i][i+1] = s[i] == s[i+1] ? 0 : 1


	res = fx[0][n-1]

	runtime : O(n)
*/

//var fx [][]int

func minInsertions(s string) int {
	n := len(s)

	////set visitor to limit recursion
	//if fx == nil {
	//	fx = make([][]int, n, n)
	//	for i := 0; i < n; i++ {
	//		fx[i] = make([]int, n, n)
	//		for j := 0; j < n; j++ {
	//			fx[i] = append(fx[i], -1)
	//		}
	//	}
	//}

	i, j := 0, n-1

	// base case
	if i == 0 && j == 0 {
		//fx[i][j] = 0
		return 0
	}
	if i == 0 && j == 1 {
		if s[i] == s[j] {
			//fx[i][j] = 0
			return 0
		}
		//fx[i][j] = 1
		return 1
	}

	if s[i] == s[j] {
		//if fx[i+1][j-1] == -1 {
		//
		//}
		//fx[i][j] = fx[i+1][j-1]
		return minInsertions(s[i+1 : j-1+1])
	}

	// if s[i] != s[j], we need + 1 because we are doing with next step
	//[+1] because if you want to get substring 1 -> 3 => s[1:4]
	//fx[i][j] = getMin(fx[i+1][j], fx[i][j-1])
	return getMin(minInsertions(s[i+1:j+1])+1, minInsertions(s[i:j-1+1])+1)
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
