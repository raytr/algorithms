package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{0, 1, 1, 0}))
}

func Solution(A []int) int {
	B := []int{}
	B = append(B, A...)

	B[0] = flip(B[0])
	countB := revert(B, 1)

	countA := revert(A, 0)

	if countA < countB {
		return countA
	}
	return countB
}

func revert(A []int, count int) int {
	for i := 1; i < len(A); i++ {
		if A[i-1] == A[i] {
			A[i] = flip(A[i])
			count++
		}
	}
	return count
}

func flip(a int) int {
	if a == 0 {
		return 1
	}
	return 0
}
