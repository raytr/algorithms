package main

//problem: https://leetcode.com/problems/happy-number

import "fmt"

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(7))
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {

	arr := separateToDigits(n)
	sumResults := make(map[int]struct{})

	for {
		sum := 0
		for _, digit := range arr {
			sum = sum + digit*digit
		}
		if sum == 1 {
			return true
		}
		if _, ok := sumResults[sum]; ok {
			return false
		}
		sumResults[sum] = struct{}{}

		arr = separateToDigits(sum)
	}

	return false
}

// create a array contain all of digits
func separateToDigits(n int) []int {
	arr := make([]int, 0)
	for n != 0 {
		if residual := n % 10; residual != 0 {
			arr = append(arr, residual)
			n = n - residual
		}

		n = n / 10
	}
	return arr
}
