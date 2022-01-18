package find_anagram_mappings

/*
	problem: https://leetcode.com/problems/find-anagram-mappings/?fbclid=IwAR0JWRWq-iq_ylI_tdTy0BsOAqRm4W_Bsh_KbjIPHx5B9tP-R5sIfd88-BM

	num2Map with key = num's value, value is index of number in num2
	loop over num1, in each iterate map[num] add to arr

	complexity: N(length of nums1)
*/

func anagramMappings(nums1 []int, nums2 []int) []int {
	result := make([]int, 0, len(nums1))
	num2Map := make(map[int]int)
	for i, v := range nums2 {
		num2Map[v] = i
	}
	for _, n := range nums1 {
		result = append(result, num2Map[n])
	}
	return result
}
