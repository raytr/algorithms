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

//func subarraysWithKDistinct(nums []int, k int) int {
//	s, f := 1, 1
//	freqMap := make(map[int]int)
//	freqMap[nums[0]]++
//	count := 0
//	for f < len(nums) {
//		freqMap[nums[f]]++
//		if len(freqMap) == k {
//			count++
//		} else if len(freqMap) > k {
//			// 1 2 1 2 3
//			//s++
//			for s < f-1 { //because not sastisfied condition at f
//				count += f - 1 - s
//				s++
//			}
//			count++ //for f, s==f-1
//			s++
//		}
//		f++
//	}
//
//	return count
//}

func subarraysWithKDistinct(nums []int, k int) int {
	/*
			2 pointers: f & s at index 0
		    freqMap = map[int]int

		    f move as posible, with each step, add subarray to res
		        while len(freqMap) < k => map[f]++, f++ (find satisfied case)
				while len(freqMap) == k => count++; map[f]++; f++
		        f--; delete map[f] (because cur f is not satisfied)
		        while len(freqMap) == k => count++; descreaseMap; s++

		    f++

		    complexity: O(logn)
	*/

	/*
		sastified =>
	*/

	s, f := 0, 0
	freqMap := make(map[int]int)
	//freqMap[nums[0]]++
	count := 0
	for f < len(nums) {
		for len(freqMap) < k {
			freqMap[nums[f]]++
			f++
		}
		//descreaseMap(freqMap, nums[f])
		for len(freqMap) == k {
			count++
			freqMap[nums[f]]++
			f++
		}

		//because cur f is not satisfied
		f--

		for len(freqMap) == k {
			count++
			//count += f - s + 1
			descreaseMap(freqMap, nums[s])
			s++
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
