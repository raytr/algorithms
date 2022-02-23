package sqrt_x

// problem: https://leetcode.com/problems/sqrtx/

func mySqrt(x int) int {
	lo := 0
	hi := x
	res := 0
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if mid*mid <= x {
			res = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return res
}
