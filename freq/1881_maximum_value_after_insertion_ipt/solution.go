package maximum_value_after_insertion

import (
	"fmt"
	"strconv"
)

/*
	problem: https://leetcode.com/problems/maximum-value-after-insertion/
*/

func maxValue(n string, x int) string {
	nVal, _ := strconv.Atoi(n)
	res := ""

	if nVal == 0 {
		res = fmt.Sprintf("%v0", x)
	} else if nVal > 0 {
		for i := 0; i < len(n); i++ {
			curV, _ := strconv.Atoi(string(n[i]))
			if curV < x {
				res = n[0:i] + strconv.Itoa(x) + n[i:]
				return res
			}
		}
		res = n + strconv.Itoa(x)
	} else {
		for i := 1; i < len(n); i++ {
			curV, _ := strconv.Atoi(string(n[i]))
			if curV > x {
				res = n[0:i] + strconv.Itoa(x) + n[i:]
				return res
			}
		}
		res = n + strconv.Itoa(x)
	}
	return res
}

// 9 8 7 6
