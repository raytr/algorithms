//package two_sum
package main

import "fmt"

//problem: https://leetcode.com/problems/two-sum/

func main() {
	fmt.Println(twoSum([]int{3, 2, 3}, 6))
}

//brute force O(n^2)
func twoSumFruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

/*
	q: is array has negative number

	theRestMap
	iterate nums
		check theRestMap[target - num] is exist ;
			if not => theRestMap[target - num] = index
			if exist => [theRestMap[target - num], i]

*/
// complexity: O(n)
func twoSum(nums []int, target int) []int {
	theRestMap := make(map[int]int)
	for i, n := range nums {
		if val, exist := theRestMap[n]; exist {
			return []int{val, i}
		} else {
			theRestMap[target-n] = i
		}
	}
	return []int{}
}
