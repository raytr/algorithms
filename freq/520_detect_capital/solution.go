package detect_capital

/*
problem: https://leetcode.com/problems/detect-capital/

check the first character, if c <= 'Z'
if second char <= 'Z' = > the rest of characters <= 'Z' = > true
else second char > 'Z' => the rest of charcters > 'Z' => true
return false

runtime: O(n)

*/
func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	if word[0] <= 'Z' {
		if word[1] <= 'Z' {
			for i := 2; i < len(word); i++ {
				if word[i] > 'Z' {
					return false
				}
			}
			return true
		} else {
			for i := 2; i < len(word); i++ {
				if word[i] <= 'Z' {
					return false
				}
			}
			return true
		}
	} else {
		for i := 1; i < len(word); i++ {
			if word[i] <= 'Z' {
				return false
			}
		}
		return true
	}
	return false
}
