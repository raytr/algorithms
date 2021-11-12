package subarrays_with_k_different_integers

/*
	problem: https://leetcode.com/problems/subarrays-with-k-different-integers

	2 pointers: f & s at index 1,0

    f increase
		while len(freqMap) == k && s > f
			count++
			s++
			decreaseMap



    complexity: O(logn)
*/

func subarraysWithKDistinct(nums []int, k int) int {
	s, f := 1, 1
	freqMap := make(map[int]int)
	freqMap[nums[0]]++
	count := 0
	for f < len(nums) {
		freqMap[nums[f]]++
		if len(freqMap) == k {
			count++
		} else if len(freqMap) > k {
			// 1 2 1 2 3
			//s++
			for s < f-1 { //because not sastisfied condition at f
				count += f - 1 - s
				s++
			}
			count++ //for f, s==f-1
			s++
		}
		f++
	}

	return count
}
