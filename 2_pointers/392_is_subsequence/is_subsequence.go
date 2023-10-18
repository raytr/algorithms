package is_subsequence

/*
	problem: https://leetcode.com/problems/is-subsequence

	t = "a h b g d c"
	s = "a b c"

	2 pointer tP and sP
	integrate tP from left to right. With each loop, we compare it's value to sP's value
	if 2 value is the same => moving sP by 1
	if sP == len(s) => return true
	otherwise => return false

	time complexity: O(N); N is length of t
	space complexity: O(1).
*/

func isSubsequence(s string, t string) bool {
	sP := 0

	for tP := 0; tP < len(t); tP++ {
		if sP < len(s) && t[tP] == s[sP] {
			sP++
		}
	}

	if sP == len(s) {
		return true
	}
	return false
}
