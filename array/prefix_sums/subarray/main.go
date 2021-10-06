package main

import (
	"fmt"
)

//Given array A of integers and an integer S, count number of subarray (continuous) that has sum == S.
//Example:
//A = [5,-1,3,2,9,4]
//S = 5
//
//=> result = 2
//There are 2 subarray sum == 5: [5], [3,2]

func main() {
	fmt.Println(naive([]int{5, -1, 3, 2, 9, 4}, 5))
}

//Complexity: O(n)
func naive(nums []int, s int) int {
	count := 0
	prefixSums := []int{}
	sumMap := make(map[int]struct{})

	//with the first element
	prefixSums = append(prefixSums, nums[0])
	if nums[0] == s {
		count++
	}

	for i := 1; i < len(nums); i++ {
		prefixSums = append(prefixSums, prefixSums[i-1]+nums[i])
		if prefixSums[i] == s {
			count++
		} else if _, exist := sumMap[prefixSums[i]]; exist {
			count++
		} else {
			sumMap[prefixSums[i]+s] = struct{}{}
		}
	}
	return count
}
