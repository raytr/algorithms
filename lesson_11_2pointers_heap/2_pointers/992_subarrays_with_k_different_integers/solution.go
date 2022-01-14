package subarrays_with_k_different_integers

/*
	problem: https://leetcode.com/problems/subarrays-with-k-different-integers

		    to use 2 pointers, this problem should be convert to :
				subarray with exact k = subarray <= k - subarray <= k - 1

			=> we have 2 algos are: find 2 subarray has unique number <= k & K-1

	count1 = subarraysWithMaxKDistictNums(num, k)
	count2 = subarraysWithMaxKDistictNums(num, k-1)
	return count1 - count2


	function subarraysWithMaxKDistictNums
		count := 0
		so, we have 2 pointers: f & s at index 0
				move f pointer from head to the end of array
				freqM[nums[f]]++
				while len(freqM) == K
					n = f - s + 1
					count += n*(n+1)/2
		return count


complexity: O(logn)
*/

func subarraysWithKDistinct(nums []int, k int) int {
	count1 := subarraysWithMaxKDistictNums(nums, k)
	count2 := subarraysWithMaxKDistictNums(nums, k-1)
	return count1 - count2
}

func subarraysWithMaxKDistictNums(nums []int, k int) int {
	count, s, f := 0, 0, 0
	freqM := make(map[int]int)
	for f < len(nums) {
		freqM[nums[f]]++
		if len(freqM) == k {
			n := f - s + 1
			a := n * (n + 1) / 2
			if s != 0 {
				a /= 2
			}
			count += a
			for len(freqM) == k {
				descreaseMap(freqM, nums[s])
				s++
			}
		}
		f++
	}
	return count
}

func descreaseMap(m map[int]int, key int) {
	if m[key] == 1 {
		delete(m, key)
	} else {
		m[key]--
	}
}
