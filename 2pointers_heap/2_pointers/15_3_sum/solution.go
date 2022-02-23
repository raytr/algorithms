package three_sum

import (
	"sort"
)

//problem: https://leetcode.com/problems/3sum
/*
   sorted

   get i and for x from front of nums to len(nums)-3
   in i loop:
       j is element that next i, k is last element
       while y <= z
           if sum (i,j,k) == target => add [i,j,k] to result
           if sum (i,j,k) > target => y--
           if sum (i,j,k) < target => z++
           .......

	complexity: O(n^2)
*/

func threeSum(nums []int) [][]int {

	res := make([][]int, 0)
	sort.Ints(nums)
	i := 0
	n := len(nums) - 1
	if len(nums) < 3 {
		return res
	}
	for i <= n-2 {
		if i == 0 || nums[i] != nums[i-1] {
			j := i + 1
			k := n
			for j < k {
				if j <= n-1 && nums[j] == nums[j-1] && j-1 > i {
					j++
					continue
				}
				if k+1 <= n && nums[k] == nums[k+1] {
					k--
					continue
				}
				sum := nums[i] + nums[j] + nums[k]
				if sum == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
					j++
					k--
				} else if sum > 0 {
					k--
				} else {
					j++
				}
			}
		}
		i++
	}
	return res
}

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)

	n := len(nums)
	res := make([][]int, 0)
	a := 0
	for a < n-2 {
		if a == 0 || nums[a] != nums[a-1] {
			b := a + 1
			c := n - 1
			for b < c {
				if b < n-2 && nums[b] == nums[b-1] && b-1 > a {
					b++
					continue
				}
				if c+1 < n && nums[c] == nums[c+1] {
					c--
					continue
				}
				sum := nums[a] + nums[b] + nums[c]
				if sum == 0 {
					res = append(res, []int{nums[a], nums[b], nums[c]})
					b++
					c--
				} else if sum < 0 {
					b++
				} else {
					c--
				}
			}
		}
		a++
	}
	return res
}
