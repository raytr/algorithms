package group_anagrams

import "sort"

/*
	https://leetcode.com/problems/group-anagrams

	integrate all elems in this string
	with each word:
		- sort it and put it to map with key is the length of this word

	n is maximum length of strs
	m is maximum characters that one element can have

	time complexity: O(n * m * logm)
	space complexity: O(n * m)
*/

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	sorted := make([]string, 0, len(strs))

	//n is maximum length of strs
	//m is maximum characters that a element can has

	// sort all character of element in strs
	for _, v := range strs { //O(n)
		charArr := make([]int32, 0, len(v)) //space: O(m)
		for _, c := range v {
			charArr = append(charArr, c)
		}
		sort.Slice(charArr, func(i, j int) bool { //O(m*logm)
			return charArr[i] < charArr[j]
		})

		sorted = append(sorted, string(charArr)) //space: O(n)
	}

	wordMap := make(map[string][]string) //space: O(n)
	for i, v := range sorted {           //O(n)
		if _, exist := wordMap[v]; !exist {
			wordMap[v] = make([]string, 0)
		}
		wordMap[v] = append(wordMap[v], strs[i])
	}

	for _, v := range wordMap {
		result = append(result, v)
	}

	return result

	//time complexity: O(n * m * logm)
	//space complexity: O(n * m)
}
