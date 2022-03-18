package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{5, 3, 1, 3, 2, 3}))
}

var existM map[string]int

func Solution(A []int) int {
	existM = make(map[string]int)
	freqSum := make(map[int]int)
	for i := 0; i < len(A)-1; i++ {
		if isValid([]int{A[i], A[i+1]}) {
			key := fmt.Sprintf("%v,%v", A[0], A[1])
			existM[key]++
			sum := A[i] + A[i+1]
			freqSum[sum]++
		}
	}

	vMax, kMax := 0, 0
	for k, v := range freqSum {
		if v > vMax {
			vMax = v
			kMax = k
		}
	}

	return kMax
}

func isValid(v []int) bool {
	k1 := fmt.Sprintf("%v,%v", v[0], v[1])
	k2 := fmt.Sprintf("%v,%v", v[1], v[0])
	if _, exist := existM[k1]; exist {
		return false
	}
	if _, exist := existM[k2]; exist {
		return false
	}
	return true
}
