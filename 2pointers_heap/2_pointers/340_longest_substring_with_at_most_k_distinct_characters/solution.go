package longest_substring_with_at_most_k_distinct_characters

/*
			problem: https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/

	   slow, fast, max = 0,0,0
	   seqMap= map[string]num
	   iterage f from first to last char in s

	       if map[s[fast]] is exist
	           map[s[fast]]++
	           update max
	           fast++
	       if len(disctinMap) == k && !map[s[fast]]
	           if map[s[slow]] == 1 => remove map[s[fast]]
	           else => map[s[fast]]--

	       return max

	       complexity: O(n)
*/

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	slow, fast, max := 0, 0, 0
	seqMap := make(map[string]int)

	for fast < len(s) && slow <= fast {
		if _, exist := seqMap[string(s[fast])]; !exist && len(seqMap) == k {
			if seqMap[string(s[slow])] == 1 {
				delete(seqMap, string(s[slow]))
			} else {
				seqMap[string(s[slow])]--
			}
			slow++
		} else {
			seqMap[string(s[fast])]++
			max = Max(max, fast-slow+1)
			fast++
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
