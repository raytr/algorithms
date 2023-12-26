package h_index

import (
	"sort"
)

/*
	problem : https://leetcode.com/problems/h-index

	[3,0,6,1,5] -> 3
	sort citations: -> 6, 5, 3, 1, 0
	interage from left to right, with each loop:
		- maxVal++
	stop meet the condition are:
		- the value < the maxVal

	reutn maxVal

	time complexity: O(N)
*/

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	index := 0

	for citations[index] > index {
		index++
		if index > len(citations)-1 {
			break
		}
	}

	return index
}
