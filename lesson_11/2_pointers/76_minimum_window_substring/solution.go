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
	/*
		   we have 2 pointers: fast & slow
		   we have 2 map: tMap & freqMap[string]index

			have 2 freqMap: tFreqMp and freqMp
			create freqMp via tArray


		   while f < len(s)
				update freqMap
		      while satisfied
					if satisfied => update min
					decreaseMap
					slow--
				fast++


	*/

	min, fast, slow := len(s)+1, 0, 0
	minSubArrIdx := []int{0, 0}
	//build tFreqMap
	tFreqMap := initTMap(t)
	freqMap := make(map[string]int)

	for fast < len(s) {
		// update freq Map
		updateFreqMap(tFreqMap, freqMap, string(s[fast]), true)

		for isSastified(tFreqMap, freqMap) {
			if fast-slow+1 < min {
				min = fast - slow + 1
				minSubArrIdx = []int{slow, fast + 1}
			}
			updateFreqMap(tFreqMap, freqMap, string(s[slow]), false)
			slow++
		}
		fast++
	}
	return s[minSubArrIdx[0]:minSubArrIdx[1]]
}

func isSastified(tMap, freqMap map[string]int) bool {
	for k, v := range tMap {
		if freqMap[k] < v {
			return false
		}
	}
	return true
}

func updateFreqMap(tMap, freqMap map[string]int, key string, increase bool) map[string]int {
	if _, exist := tMap[key]; exist {
		if increase {
			freqMap[key]++
		} else {
			if freqMap[key] == 1 {
				delete(freqMap, key)
			} else {
				freqMap[key]--
			}
		}
	}
	return freqMap
}

func initTMap(t string) map[string]int {
	tMap := make(map[string]int)
	for _, c := range t {
		tMap[string(c)]++
	}
	return tMap
}
