//package subarray_sums_divisible_by_k
package main

import "fmt"

func main() {
	//fmt.Println(subarraysDivByKByBruteForce([]int{4,5,0,-2,-3,1}, 5))
	fmt.Println(subarraysDivByKByBruteForce([]int{-5}, 5))
}

// problem: https://leetcode.com/problems/subsets/?fbclid=IwAR36lGWQEjXiAGh8ZNw0qrl5BrXlQdAxjyBchgdG1s9O7EA86Rp6vktJCxw

//complexity : O(n*2)
func subarraysDivByKByBruteForce(nums []int, k int) int {
	if len(nums) == 1 && nums[0]%k == 0 {
		return 1
	}

	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]%k == 0 {
			count++
		}
		sum := 0
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			a := nums[i] + sum
			if a%k == 0 {
				count++
			}
		}
	}
	return count
}

/*
	loop over the array, go check which number can mod k == 0 => if has => count++
	calculate prefix sum: prifixSum[i] = prefixSum[i-1]+sums[i]

	have a map to store all prefix sums that checked

	check all prefix sum,
		if prefix sum % k == 0 => if has => count++
		else if prefix sum exist in map => add to map
		else
*/
func subarraysDivByK(nums []int, k int) int {
	count := 0
	for _, n := range nums {
		if n%k == 0 {
			count++
		}
	}

	//4,9,9,7,4,5

	//calculate prefix sum
	prefixSums := make([]int, 0, len(nums))
	prefixSums = append(prefixSums, nums[0])
	for i := 1; i < len(nums); i++ {
		prefixSums = append(prefixSums, prefixSums[i-1], nums[i])
	}

	for i := 1; i < len(nums); i++ {
		if prefixSums[i]%k == 0 {
			count++
		} else {

		}
	}

	return count
}
