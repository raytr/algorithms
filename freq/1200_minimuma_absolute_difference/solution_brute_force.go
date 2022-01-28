package minimuma_absolute_difference

import "math"

/*
   set min is max int
   we iterate every possible pair in the array
   if abs(a,b) == min appent it to result
   else if abs(a,b) < min => clear all element in res and append pair of a,b

   time complexity is O(n^2)

	problem: https://leetcode.com/problems/minimum-absolute-difference/
*/

func minimumAbsDifferenceBruteForce(arr []int) [][]int {
	n := len(arr)
	min := math.Inf(1)
	res := make([][]int, 0, n)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			absValue := math.Abs(float64(arr[i] - arr[j]))
			if absValue < min {
				min = absValue
				res = make([][]int, 0, n)
				res = append(res, []int{arr[i], arr[j]})
			} else if absValue == min {
				res = append(res, []int{arr[i], arr[j]})
			}
		}
	}
	return res
}
