package subarray_sum_equal_k

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
