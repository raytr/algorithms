package anagrams

import "sort"

/*
	interay all elem in this string
	with each word:
		- sort it and put it to map with key is the lenge of

	N is the lenght of array
	M is the length of each element


*/

func anagrams(arr []string) [][]string {

	result := make([][]string, 0)
	sorted := make([]string, 0, len(arr))

	// sort all character of element in arr
	for _, v := range arr {
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
		wordMap[v] = append(wordMap[v], arr[i])
	}

	for _, v := range wordMap {
		result = append(result, v)
	}

	return result
}
