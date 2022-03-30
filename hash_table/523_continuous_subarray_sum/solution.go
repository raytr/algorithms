package continuous_subarray_sum

/*
	problem: https://leetcode.com/problems/continuous-subarray-sum/

	- convert to prefix sum by formula: prefixSum[i] = (prefixSum[i-1] + nums[i]) % k
	- create a map : prefixSumMap with key = mod, value = frequently of mod
	- if practice > 2 or practice == 2 && prefixSum[i-1] != nums[i]
	because practice == 2 && prefixSum[i-1] == nums[i] happens when just has 1 number mod k == 0
	=> subarray has only 1 element

	complexity: O(n)
*/

func checkSubarraySum(nums []int, k int) bool {
	prefixSums := []int{nums[0] % k}
	freq := make(map[int]int)
	freq[0]++
	freq[nums[0]%k]++
	//change to prefix Sum
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 && nums[i-1] == 0 {
			return true
		}
		prefixSums = append(prefixSums, (prefixSums[i-1]+nums[i])%k)
		freq[prefixSums[i]]++
		if freq[prefixSums[i]] > 2 || (freq[prefixSums[i]] == 2 && prefixSums[i] != prefixSums[i-1]) {
			return true
		}
	}
	return false
}
