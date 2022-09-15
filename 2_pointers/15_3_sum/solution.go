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
			l := i + 1
			r := n
			for l < r {
				if l <= n-1 && nums[l] == nums[l-1] && l-1 > i {
					l++
					continue
				}
				if r+1 <= n && nums[r] == nums[r+1] {
					r--
					continue
				}
				sum := nums[i] + nums[l] + nums[r]
				if sum == 0 {
					res = append(res, []int{nums[i], nums[l], nums[r]})
					l++
					r--
				} else if sum > 0 {
					r--
				} else {
					l++
				}
			}
		}
		i++
	}
	return res
}

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0, 0)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		//avoid duplicate at i index
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := n - 1
		for l < r {
			//avoid duplicate at l & r index
			if l-1 > i && l < n-1 && nums[l] == nums[l-1] && r > l {
				l++
				continue
			}
			if r < n-2 && nums[r] == nums[r+1] {
				r--
				continue
			}

			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
