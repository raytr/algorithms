package valid_anagram

/*
	problem: https://leetcode.com/problems/valid-anagram


	time complexity: O(M + N) M is length of s and N is length of t
	space complexity: O(M+N)
*/

func isAnagram(s string, t string) bool {

	if len(s) == 0 || len(t) == 0 || len(s) != len(t) {
		return false
	}

	sFreqMap := make(map[int32]int)
	tFreqMap := make(map[int32]int)

	for _, c := range s {
		sFreqMap[c]++
	}

	for _, c := range t {
		tFreqMap[c]++
	}

	for sKey, sFreq := range sFreqMap {
		if tFreqMap[sKey] != sFreq {
			return false
		}
	}

	for tKey, tFreq := range tFreqMap {
		if sFreqMap[tKey] != tFreq {
			return false
		}
	}

	return true
}
