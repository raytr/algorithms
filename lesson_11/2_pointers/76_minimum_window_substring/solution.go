package minimum_window_substring

/*
	problem: https://leetcode.com/problems/minimum-window-substring

   s, min=0, len(s)
   minSubArray := [0, len(s)-1]
   tSeqMap := map[string]int to store all frequency of characters in t

   for f=0, f<len(s);f++
       if !tMap[s[f]]
           if s==f => s++, continue
       else
           tSeqMap[s[f]]++




   return minSubArray
   complexity: O(f+s)
*/

func minWindow(s string, t string) string {
	slow, min := 0, len(s)-1
	minSubArray := s
	tMap := initTMap(t)

	for fast := 0; fast < len(s); fast++ {
		if _, exist := tMap[string(s[fast])]; !exist {
			if slow == fast {
				slow++
				continue
			}
		} else {
			delete(tMap, string(s[fast]))
			if len(tMap) == 0 {
				min = Min(min, fast-slow+1)
				minSubArray = s[slow : fast+1]
			}
		}
	}
	return minSubArray
}

func initTMap(t string) map[string]int {
	tMap := make(map[string]int)
	for _, c := range t {
		tMap[string(c)]++
	}
	return tMap
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
