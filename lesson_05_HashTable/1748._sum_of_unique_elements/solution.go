//package sum_of_unique_elements
package main

import "fmt"

func main() {
	fmt.Println(sumOfUnique([]int{1, 2, 3, 2}))
	fmt.Println(sumOfUnique([]int{1, 1, 1, 1, 1}))
	fmt.Println(sumOfUnique([]int{1, 2, 3, 4, 5}))
}

/*
	problem: https://leetcode.com/problems/sum-of-unique-elements/?fbclid=IwAR29bn7vJoBUQ3y1IRcJa_SGYShO7lh1ElQjHODaht4Yi6ro7y1cqiQiO4M

	create a map with key is number and value is the frequency of this number
	loop over the s => add to this number's frequency
	loop over numberMap, if value == 1 => sum += key
	return sum

	complexity: O(length of nums)
*/

func sumOfUnique(nums []int) int {
	numberMap := make(map[int]int)
	for _, num := range nums {
		numberMap[num]++
	}
	sum := 0
	for k, v := range numberMap {
		if v == 1 {
			sum += k
		}
	}
	return sum
}
