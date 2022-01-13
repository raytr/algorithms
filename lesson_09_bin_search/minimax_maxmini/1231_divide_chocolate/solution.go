package divide_chocolate

/*
	problem: https://leetcode.com/problems/divide-chocolate/

	you eat minimum total sweetness find max sweetness you can get in => maxmini

	hi = sum(sweetness)
	lo = 1

	bin => while hi >= lo
	if maxShared < k + 1 => not satisfied => increase maxShared => decrease mid => hi = mid - 1
	if maxShared >= k + 1 => satisfied => res = mid; let check bigger mid => lo = mid + 1

	complexity: O(n*logn)
*/

func maximizeSweetness(sweetness []int, k int) int {
	var lo, hi, res int

	for _, s := range sweetness {
		hi += s
	}

	for lo <= hi {
		mid := lo + (hi-lo)/2
		var maxShared, piece int

		//calculate maxShared
		for _, s := range sweetness {
			piece += s
			if piece >= mid {
				maxShared++
				piece = 0
			}
		}

		if maxShared < k+1 {
			hi = mid - 1
		} else {
			res = mid
			lo = mid + 1
		}
	}
	return res
}
