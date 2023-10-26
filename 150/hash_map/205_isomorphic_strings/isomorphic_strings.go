package isomorphic_strings

/*
	problem: https://leetcode.com/problems/isomorphic-strings/

	create a map with key is character of s, value is character of t
	integrate all characters in s,
		if it doesn't exist , new key value
		otherwise, if map[character in s] has value that not same with the character in t, return false

	do the same with t

	time complexity : O(N), N is the number of characters in the longest string
*/

func isIsomorphic(s string, t string) bool {
	sMap := make(map[uint8]uint8)
	tMap := make(map[uint8]uint8)

	for i := 0; i < len(s); i++ {
		if _, exist := sMap[s[i]]; !exist {
			sMap[s[i]] = t[i]
		} else if sMap[s[i]] != t[i] {
			return false
		}

		if _, exist := tMap[t[i]]; !exist {
			tMap[t[i]] = s[i]
		} else if tMap[t[i]] != s[i] {
			return false
		}
	}

	return true
}
