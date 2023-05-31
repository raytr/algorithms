package second

import (
	"strconv"
)

func TheLargestNumber(N int) int {
	nStr := strconv.Itoa(N)
	n := len(nStr)
	resStr := ""
	if N == 0 {
		resStr = "50"
	} else if N < 0 {
		for i := 1; i < n; i++ {
			if nStr[i] > '5' {
				resStr = nStr[0:i] + "5" + nStr[i:]
				break
			}
		}
	} else {
		for i := 0; i < n; i++ {
			if nStr[i] < '5' {
				resStr = nStr[0:i] + "5" + nStr[i:]
				break
			}
		}
	}

	res, _ := strconv.Atoi(resStr)
	return res
}
