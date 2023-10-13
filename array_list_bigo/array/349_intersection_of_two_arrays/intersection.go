package intersection_of_two_arrays

/*
	https://leetcode.com/problems/intersection-of-two-arrays/

*/

func intersection(nums1 []int, nums2 []int) []int {

	//time complexity: O(N). N is total elements of 2 array

	nums1Map := make(map[int]bool)
	nums2Map := make(map[int]bool)
	result := make([]int, 0)

	for _, n := range nums1 {
		nums1Map[n] = true
	}

	for _, n := range nums2 {
		nums2Map[n] = true
	}

	for k, _ := range nums1Map {
		if _, exist := nums2Map[k]; exist {
			result = append(result, k)
		}
	}

	return result
}
