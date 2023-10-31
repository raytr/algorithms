package word_pattern

import "strings"

/*
	problem: https://leetcode.com/problems/word-pattern
*/

func wordPattern(pattern string, s string) bool {
	arrs := strings.Split(s, " ")
	if len(arrs) != len(pattern) {
		return false
	}

	patternSPair := make(map[uint8]string)
	sPatternPair := make(map[string]uint8)

	for i := 0; i < len(pattern); i++ {
		if _, exist := patternSPair[pattern[i]]; exist {
			if patternSPair[pattern[i]] != arrs[i] {
				return false
			}
		} else {
			patternSPair[pattern[i]] = arrs[i]
		}

		if _, exist := sPatternPair[arrs[i]]; exist {
			if sPatternPair[arrs[i]] != pattern[i] {
				return false
			}
		} else {
			sPatternPair[arrs[i]] = pattern[i]
		}
	}

	return true
}
