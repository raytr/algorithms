package minimum_window_substring

/*
	problem: https://leetcode.com/problems/minimum-window-substring

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



ATTENTION:
"improve complexity: creare freqMap by negative number. When meet a number that is contained in freqMap
=> ++ => if freqMap[key] == 0 => count++
=> satisfied when count == uniqueNumber"


	func satisfied() {
		if count == uniqueNumberTMap => return true
		else return false
}


   return minSubArray
   complexity: O(f+s)
*/

var count, uniqueTMap = 0, 0

func minWindow(s string, t string) string {
	count = 0
	min, fast, slow := len(s)+1, 0, 0
	minSubArrIdx := []int{0, 0}

	freqMap := initFreqMap(t)
	uniqueTMap = len(freqMap)

	for fast < len(s) {
		// update practice Map
		updateFreqMap(freqMap, string(s[fast]), true)

		for isSastified() {
			if fast-slow+1 < min {
				min = fast - slow + 1
				minSubArrIdx = []int{slow, fast + 1}
			}
			updateFreqMap(freqMap, string(s[slow]), false)
			slow++
		}
		fast++
	}
	return s[minSubArrIdx[0]:minSubArrIdx[1]]
}

func isSastified() bool {
	if count == uniqueTMap {
		return true
	}
	return false
}

func updateFreqMap(freqMap map[string]int, key string, increase bool) map[string]int {
	if _, exist := freqMap[key]; exist {
		if increase {
			freqMap[key]++
			if freqMap[key] == 0 {
				count++
			}
		} else {
			if freqMap[key] == 0 {
				count--
			}
			freqMap[key]--
		}
	}
	return freqMap
}

func initFreqMap(t string) map[string]int {
	tMap := make(map[string]int)
	for _, c := range t {
		tMap[string(c)]--
	}
	return tMap
}
