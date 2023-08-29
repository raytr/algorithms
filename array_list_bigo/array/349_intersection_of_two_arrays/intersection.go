package intersection_of_two_arrays

/*
	https://leetcode.com/problems/intersection-of-two-arrays/
*/

func intersection(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]bool)
	resultMap := make(map[int]bool)
	result := make([]int, 0)

	for _, n := range nums1 {
		nums1Map[n] = true //4 ,9 ,5
	}

	for _, num := range nums2 {
		if _, ok := nums1Map[num]; ok {
			resultMap[num] = true // 9, 4,
		}
	}

	for k, _ := range resultMap {
		result = append(result, k)
	}

	return result
}
