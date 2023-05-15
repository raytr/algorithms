package consecutive_characters

// https://leetcode.com/problems/consecutive-characters

func maxPower(s string) int {
	count := 1
	max := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}
		max = getMax(max, count)
	}
	return max
}

func getMax(a, b int) int {
	if a < b {
		return a
	}
	return b
}
