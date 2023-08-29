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

//brute force: with each number, get sum from i + 1 to end, when sum == k => count++

func subarraySumBruteForce(nums []int, k int) int {
	count, n := 0, len(nums)
	for start := 0; start < n; start++ {
		for end := start + 1; end <= n; end++ {
			sum := 0
			for i := start; i <= end; i++ {
				sum += nums[i]
			}
			if sum == k {
				count++
			}
		}
	}
	return count
}
