package contains_duplicate_ii

/*
	valueMap : map[value]index
	iterate over the nums
	if map[v] is exist => i - map[v] <= k => true
	false
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	appearsAt := make(map[int]int)
	for i, v := range nums {
		a := i - appearsAt[v]
		b := a <= k
		if _, exist := appearsAt[v]; exist && b {
			return true
		}
		appearsAt[v] = i
	}
	return false
}
