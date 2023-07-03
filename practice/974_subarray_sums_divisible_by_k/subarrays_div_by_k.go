package subarray_sums_divisible_by_k

/*
	problem: https://leetcode.com/problems/longest-consecutive-sequence/submissions/981305698/
	-10000 < nums[i] < 10000


			h
	nums = [4,5,0,-2,-3,1], k = 5
			  t

	check is can tail divisible by k
	check sum of subarray from h to t can divisible by k

	time complexity is O(n^3)

1,2,3
1

*/

func subarraysDivByKBruteForce(nums []int, k int) int {
	count := 0
	n := len(nums)

	for h := 0; h < n; h++ {
		for t := h; t < n; t++ {
			if h == t && isDivisible(nums[t], k) { //1 -> count = 0;  2  -> count = 1;  3  -> count = 1
				count++
			}
			if t > h {
				sum := 0
				for i := h; i <= t; i++ {
					sum += nums[i]
				}
				if isDivisible(sum, k) { // sum = 3 -> count = 1; sum = 6  -> count = 2
					count++
				}
			}
		}
	}

	return count
}

/*

	class Solution {
    public int subarraysDivByK(int[] nums, int k) {
        int n = nums.length;
        int prefixMod = 0, result = 0;

        // There are k mod groups 0...k-1.
        int[] modGroups = new int[k];
        modGroups[0] = 1;

        for (int num: nums) {
            // Take modulo twice to avoid negative remainders.
            prefixMod = (prefixMod + num % k + k) % k;
            // Add the count of subarrays that have the same remainder as the current
            // one to cancel out the remainders.
            result += modGroups[prefixMod];
            modGroups[prefixMod]++;
        }

        return result;
    }
}

*/

func subarraysDivByK(nums []int, k int) int {
	//n := len(nums)
	prefixMod := 0
	result := 0

	// There are k mod groups 0...k-1.
	modGroups := make([]int, k, k)
	modGroups[0] = 1

	for _, num := range nums {
		// Take modulo twice to avoid negative remainders.
		prefixMod = (prefixMod + num%k + k) % k
		// Add the count of subarrays that have the same remainder as the current
		// one to cancel out the remainders.
		result += modGroups[prefixMod]
		modGroups[prefixMod]++
	}

	return result

}

func isDivisible(num, k int) bool {
	if num%k == 0 {
		return true
	}
	return false
}
