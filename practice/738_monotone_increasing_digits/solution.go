// package monotone_increasing_digits
package main

import "fmt"

/*
	https://leetcode.com/problems/monotone-increasing-digits/
*/

func main() {
	fmt.Println(monotoneIncreasingDigits(322))
}

func monotoneIncreasingDigits(n int) int {
	if n < 10 {
		return n
	}

	//check if n is 10,100,1000,...
	if n%10 == 0 {
		return n - 1
	}

	digist := make([]int, 0)
	num := n

	//put all digits to digits
	for num >= 1 {
		digist = append(digist, num%10)
		num = num / 10
	}

	digist = revertSlice(digist)

	//check num is monotone increaseing digits
	isMono := true
	for i := 0; i < len(digist)-1; i++ {
		if digist[i] <= digist[i+1] {
			continue
		} else {
			isMono = false
			break
		}
	}
	if isMono {
		return n
	}

	for i := len(digist) - 1; i > 0; i-- {
		if digist[i-1] > digist[i] {
			digist[i-1] = digist[i] - 1
			digist[i] = 9
		}
	}

	//convert to number
	result := 0
	for i := 0; i < len(digist); i++ {
		result += digist[i] * (i + 1)
	}

	return result
}

func revertSlice(arr []int) []int {
	n := len(arr)
	l := 0
	r := n - 1

	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}

	return arr
}
