package main

import (
	"fmt"
)

func findMaxSum(numbers []int) int {
	n := len(numbers)
	a, b := 0, n-1
	max := -100000
	maxA, maxB := numbers[a], numbers[b]
	for i, num := range numbers {
		if num > numbers[a] {
			a = i
			b = a
		} else if num > numbers[b] {
			b = i
		}
	}

	for a < b {
		if num > numbers[a] {
			a = i
			b = a
		} else if num > numbers[b] {
			b = i
		}
	}

	return numbers[a] + numbers[b]
}

func main() {
	list := []int{5, 9, 7, 11}
	fmt.Println(findMaxSum(list))
}
