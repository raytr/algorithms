package sum_closest

import (
	"math"
	"sort"
)

/*

		https://leetcode.com/problems/3sum-closest/

	   if len == 3 => return sum

		first, sort nums
		ex: -5 -5 -4 0 0 3 3 4 5
		     p  h              t

		we have 3 pointers: pivot, l, r
	   pilot, l = pilot + 1, r = the end of array
		if sum of them > target => move r to left, oderwise : move l to right, if sum == target, return 0
		move pivot by one and make the similar process until pivot == len of nums - 2

		complexity: O(N^2)

*/

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)

	minSum := 3000
	minDis := 30000
	for p := 0; p <= n-2; p++ {
		l := p + 1
		r := n - 1

		for l < r {
			sum := nums[p] + nums[l] + nums[r]
			if sum == target {
				return sum
			}

			dis := int(math.Abs(float64(sum - target)))
			if dis < minDis {
				minDis = dis
				minSum = sum
			}
			if sum < target {
				l++
			} else {
				r--
			}
		}
	}

	return minSum
}
