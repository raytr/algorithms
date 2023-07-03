package minimuma_absolute_difference

import (
	"sort"
)

/*
	problem: https://leetcode.com/problems/minimum-absolute-difference/


	sort arr first
	integrate all element in arr to check i with i+1, if arr[i+1] - a[i] < min
	   - make this result empty
	   - update min
   else if arr[i] == arr[i+1] => add to result

	run time complexity O(nLogn)
*/

func minimumAbsDifferenceSortingFirst(arr []int) [][]int {
	n := len(arr)
	min := 100000
	result := make([][]int, 0, n)

	sort.Ints(arr)

	for i := 0; i < n-1; i++ {
		diff := (arr[i+1] - arr[i]) * (arr[i+1] - arr[i]) / (arr[i+1] - arr[i])
		if diff < min {
			result = make([][]int, 0, n)
			result = append(result, []int{arr[i], arr[i+1]})
			min = diff
		} else if diff == min {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}
	return result
}
