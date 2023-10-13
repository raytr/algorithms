package intersection_of_two_arrays_ii

/*
	problem: https://leetcode.com/problems/intersection-of-two-arrays-ii/

	HOW SHOULD I HANDLER DUPLICATE?
	what is range of length of nums1 and 2
	what is minimum and maximum of each element in nums1&nums2
	what happen if nums1 & nums2 doesn't has any intersection

*/

func intersect(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)

	freqMapNums1 := make(map[int]int)
	for _, n := range nums1 {
		freqMapNums1[n]++
	}

	for _, n := range nums2 {
		if freqMapNums1[n] > 0 {
			result = append(result, n)
			freqMapNums1[n]--
		}
	}

	return result
}
