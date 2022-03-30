package minimuma_absolute_difference

import (
	"math"
	"sort"
)

/*
   set min is max int


	1 4 6 7 9 11 14 15
	 3 2 1 2 2  3  1

	minAbsDiff = max of int
	we check pair of 2 numbers next to each other,
		if curAbsDiff < minAbsDiff => discard the res and add pair of current nums
		else if curAbsDiff == minAbsDiff => add 2 cur nums to the res

	return res
*/

func minimumAbsDifferenceOptiomized(arr []int) [][]int {
	sort.Ints(arr)
	n, minAbsDiff := len(arr), math.Inf(1)
	res := make([][]int, 0, n)
	for i := 1; i < n; i++ {
		curAbsDiff := math.Abs(float64(arr[i] - arr[i-1]))
		if curAbsDiff < minAbsDiff {
			minAbsDiff = curAbsDiff
			res = make([][]int, 0, n)
			res = append(res, []int{arr[i-1], arr[i]})
		} else if curAbsDiff == minAbsDiff {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}
