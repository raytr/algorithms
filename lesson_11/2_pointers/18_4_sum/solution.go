package four_sum

import (
	"sort"
)

//problem: https://leetcode.com/problems/4sum

/*
   sort

   range a from 0 to len-4
       range b from a+1 to len-3
           c= b+1
           d= len-1
           while c < d
               if sum == target => add [a,b,c,d] to res
               if sum < target => c++
               if sum > target => d--


               0 1 2 3 4 5 6
*/

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	a, n := 0, len(nums)-1

	if len(nums) < 4 {
		return res
	}

	sort.Ints(nums)

	for a <= n-3 {
		if a == 0 || nums[a] != nums[a-1] {
			b := a + 1
			for b <= n-2 {
				//-2,-1,-1,1,1,2,2
				if b == a+1 || nums[b] != nums[b-1] {
					c := b + 1
					d := n
					for c < d {
						if c > b+1 && c <= n-1 && nums[c] == nums[c-1] {
							c++
							continue
						}
						if d < n && nums[d] == nums[d+1] {
							d--
							continue
						}
						sum := nums[a] + nums[b] + nums[c] + nums[d]
						if sum == target {
							res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
							c++
							d--
						} else if sum < target {
							c++
						} else {
							d--
						}
					}
				}
				b++
			}
		}
		a++
	}
	return res
}
