package subarray_sum_equal_k

/*
	problem: https://leetcode.com/problems/subarray-sum-equals-k/

	add 0 at head of nums array
	convert nums to prefix sum by fomular: prefixSum[i] = prefix[i-1] + nums[i]
	for start = 0; start < n-1; start++
		for end = start + 1; end < n; end++
			if nums[end] - nums[start] == k => count++

	runtime: O(n^2)
	space: O(1)
*/

func subarraySum(nums []int, k int) int {
	//add 0 at head of nums
	nums = append([]int{0}, nums...)
	count, n := 0, len(nums)

	//convert nums to prefix sum
	for i := 1; i < n; i++ {
		nums[i] = nums[i-1] + nums[i]
	}

	for start := 0; start < n-1; start++ {
		for end := start + 1; end < n; end++ {
			if nums[end]-nums[start] == k {
				count++
			}
		}
	}
	return count
}

////Given array A of integers and an integer S, count number of 560_subarray_sum_equal_k (continuous) that has sum == S.
////Example:
////A = [5,-1,3,2,9,4]
////S = 5
////
////=> result = 2
////There are 2 560_subarray_sum_equal_k sum == 5: [5], [3,2]
//
////problem: https://leetcode.com/problems/subarray-sum-equals-k/
//
////Complexity: O(n)
//func subarraySum(nums []int, s int) int {
//	count := 0
//	prefixSums := []int{}
//	sumMap := make(map[int]struct{})
//
//	//with the first element
//	prefixSums = append(prefixSums, nums[0])
//	sumMap[nums[0]] = struct{}{}
//	if nums[0] == s {
//		count++
//	}
//
//	for i := 1; i < len(nums); i++ {
//		prefixSums = append(prefixSums, prefixSums[i-1]+nums[i])
//		if nums[i] == s {
//			count++
//		}
//		if prefixSums[i] == s {
//			count++
//		} else if _, exist := sumMap[prefixSums[i]-s]; exist {
//			count++
//		}
//		sumMap[prefixSums[i]] = struct{}{}
//	}
//	return count
//}
