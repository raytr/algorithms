package powx_n

import (
	"math"
)

/*
	problem: https://leetcode.com/problems/powx-n
*/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 1 / positivePow(x, -n)
	} else {
		return positivePow(x, n)
	}
}

func positivePow(x float64, n int) float64 {
	ori := x
	for i := 1; i < n; i++ {
		x *= ori
	}

	return math.Round(x*1000) / 1000
}
