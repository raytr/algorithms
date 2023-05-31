package _9_group_anagrams

import "sort"

/*
	https://leetcode.com/problems/group-anagrams

	integrate all elem in this string
	with each word:
		- sort it and put it to map with key is the lenge of

	N is the length of array
	K is a maximum length of the string in strs

	time complexity: O(N * KlogK)
	space complexity O(K*N)
*/

func groupAnagrams(strs []string) [][]string {

	result := make([][]string, 0)
	sorted := make([]string, 0, len(strs))

	// sort all character of element in strs
	for _, v := range strs {
		charArr := make([]int32, 0, len(v))
		for _, c := range v {
			charArr = append(charArr, c)
		}
		sort.Slice(charArr, func(i, j int) bool {
			return charArr[i] < charArr[j]
		})

		sorted = append(sorted, string(charArr))
	}

	wordMap := make(map[string][]string)
	for i, v := range sorted {
		if _, exist := wordMap[v]; !exist {
			wordMap[v] = make([]string, 0)
		}
		wordMap[v] = append(wordMap[v], strs[i])
	}

	for _, v := range wordMap {
		result = append(result, v)
	}

	return result
}
