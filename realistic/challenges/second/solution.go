package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Solution(4268))
	fmt.Println(Solution(6750))
	fmt.Println(Solution(-999))
	fmt.Println(Solution(-394))
	fmt.Println(Solution(5268))
}

func Solution(N int) int {
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

func Test() {
	text := "-394"
	r := []rune(text)
	a := append(append(append([]rune{}, r[1:2]...), '5'), r[2:]...)
	fmt.Println(string(a))
}
