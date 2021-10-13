package main

//Given array A of integers and an integer S, count number of subarray_sum_equal_k (continuous) that has sum == S.
//Example:
//A = [5,-1,3,2,9,4]
//S = 5
//
//=> result = 2
//There are 2 subarray_sum_equal_k sum == 5: [5], [3,2]

//Complexity: O(n)
func subarraySum(nums []int, s int) int {
	count := 0
	prefixSums := []int{}
	sumMap := make(map[int]struct{})

	//with the first element
	prefixSums = append(prefixSums, nums[0])
	sumMap[nums[0]] = struct{}{}
	if nums[0] == s {
		count++
	}

	for i := 1; i < len(nums); i++ {
		prefixSums = append(prefixSums, prefixSums[i-1]+nums[i])
		if nums[i] == s {
			count++
		}
		if prefixSums[i] == s {
			count++
		} else if _, exist := sumMap[prefixSums[i]-s]; exist {
			count++
		}
		sumMap[prefixSums[i]] = struct{}{}
	}
	return count
}
