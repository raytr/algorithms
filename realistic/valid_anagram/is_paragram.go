package valid_anagram

import "sort"

func isAnagram(s string, t string) bool {
	/*

		   with each string, sort all characters in this
		   if 2 string equal each other => retunr true, otherwise => return false
			N is length of s
			M is length of t
			time complexity: O((M * N)Log(M+N))
			space complexity: O(M+N)
	*/

	runeS := []rune(s)
	runeT := []rune(t)

	if len(runeS) == 0 || len(runeT) == 0 || len(runeT) != len(runeS) {
		return false
	}

	sort.Slice(runeS, func(i, j int) bool {
		return runeS[i] < runeS[j]
	})

	sort.Slice(runeT, func(i, j int) bool {
		return runeT[i] < runeT[j]
	})

	for i := 0; i < len(runeS); i++ {
		if runeS[i] != runeT[i] {
			return false
		}
	}

	return true
}