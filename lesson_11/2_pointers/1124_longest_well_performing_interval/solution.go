//package _124_longest_well_performing_interval
package main

import "fmt"

func main() {
	//fmt.Println(longestWPI([]int{9, 9, 6, 0, 6, 6, 9}))
	//fmt.Println(longestWPI([]int{6,6,6}))
	fmt.Println(longestWPI([]int{6, 9, 9}))
	//fmt.Println(longestWPI([]int{6, 6, 6, 9, 9, 6, 6}))
	//fmt.Println(longestWPI([]int{6,6,9,9,9,9,6,6}))
}

/*

	    revert array with 1 with hour > 8 and -1 for ortherwise

        iterate all of the hours to calculate prefixsum
        iterate all of prefixSums with while i
			if prxSum[i] > 0 => max++
			else if prxSum[i] == idxMap[prxSum[i]-1]
					=> m := i - prxSum[i] == idxMap[prxSum[i]-1]; find Max
					(for case that all of element prxSum subarray are negative number)
			else => idxMap[prxSum[i]-1] = i


        return max

        complexity: O(n)
*/

func longestWPI(hours []int) int {
	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
	}

	//calculate prefix sum
	for i := 1; i < len(hours); i++ {
		hours[i] = hours[i-1] + hours[i]
	}

	max := 0
	idxMap := make(map[int]int)
	for i := 0; i < len(hours); i++ {

		// for sub arr has possitivew number
		if hours[i] > 0 {
			max++
		}

		// for sub arr just has negative number
		if _, exist := idxMap[hours[i]-1]; exist {
			count := i - (idxMap[hours[i]-1])
			max = Max(max, count)
		} else {
			idxMap[hours[i]] = i
		}
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
