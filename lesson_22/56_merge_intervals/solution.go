package merge_intervals

import "sort"

/*
	problem: https://leetcode.com/problems/merge-intervals/

	min, max = intervals[0][0], intervals[0][1]
   go over this array from 1 to n - 2
       compare pair[i] and pair[i-1]
           if pair[i-1][1] < pair[i][0]{
               push to res [min, max]
               min, max = pair[i][0],pair[i][1]
           }

           else {
               find min, max and update
           }


   because each pair in intervals was travelsed once => time complexity is O(n)
   sort => O(logn)
   => total O(logn)
   and space complexity is O(1)
*/

func merge(intervals [][]int) [][]int {
	//sort based on first elemtn in pair
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	n := len(intervals)
	res := make([][]int, 0, n)
	min, max := intervals[0][0], intervals[0][1]

	for i := 1; i < n; i++ {
		if max < intervals[i][0] {
			res = append(res, []int{min, max})
			min, max = intervals[i][0], intervals[i][1]
		} else {
			if min > intervals[i][0] {
				min = intervals[i][0]
			}
			if max < intervals[i][1] {
				max = intervals[i][1]
			}
		}
	}

	//push the last pair to res
	res = append(res, []int{min, max})
	return res
}
