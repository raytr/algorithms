package find_the_smallest_divisor_given_a_threshold

/*

	problem: https://leetcode.com/problems/find-the-smallest-divisor-given-a-threshold/

   lo, ih = 1, sum(nums)

   while lo <= hi
       if isSatisfied => store res and hope smaller mid => hi = mid - 1
       else => lo = mid + 1

   isSatisfier(divisor) {
       sumDivisionRes := 0
       for num range of nums{
           val = num/divisor
           if num%divisor > 0 => val++
             sumDivisionRes += val
       }
       if sumDivisionRes > threshold => return false
       else return true
   }

*/

func smallestDivisor(nums []int, threshold int) int {
	res, lo, hi := 1, 0, 0
	for _, n := range nums {
		hi += n

	}

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if isValid(nums, mid, threshold) {
			res = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return res
}

func isValid(nums []int, divisor, threshold int) bool {
	if divisor == 0 {
		return false
	}
	sumDivisionRes := 0
	for _, num := range nums {
		val := num / divisor
		if num%divisor > 0 {
			val++
		}
		sumDivisionRes += val
	}
	if sumDivisionRes > threshold {
		return false
	}
	return true
}
