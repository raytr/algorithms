package contains_duplicate_ii

/*
	valueMap : map[value]index
	iterate over the nums
	if map[v] is exist => i - map[v] <= k => true
	false
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	valueMap := make(map[int]int)
	for i, v := range nums {
		if _, exist := valueMap[v]; exist && i-valueMap[v] <= k {
			return true
		}
		valueMap[v] = i
	}
	return false
}
