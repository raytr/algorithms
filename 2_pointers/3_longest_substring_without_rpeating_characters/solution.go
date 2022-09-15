package longest_substring_without_rpeating_characters

/*
  problem: https://leetcode.com/problems/longest-substring-without-repeating-characters/

   max = 0
       a b c a b c b b
       s f => not exist duplicate => satisfied 2 disctin char > max => max = 2 => f++
       s   f => not exist duplicate => satisfied  => 3 disctin char > max => max = 3 => f++
       s.    f => exist duplicate => not satisfied => s++; remove existMap[string(s[slow])]
         s.  f => exist duplicate => not satisfied => s++; remove existMap[string(s[slow])]
         ......

        f == len - 1 => return max

		complexity: O(n)
*/

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	max, slow, fast := 1, 0, 1
	existMap := make(map[string]bool)
	existMap[string(s[0])] = true
	for fast < len(s) {
		// check is satisfied
		if _, exist := existMap[string(s[fast])]; !exist {
			max = Max(max, fast-slow+1)
			existMap[string(s[fast])] = true
			fast++
		} else {
			delete(existMap, string(s[slow]))
			slow++
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
