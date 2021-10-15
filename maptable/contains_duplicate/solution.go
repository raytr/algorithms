package contains_duplicate

//problem: https://leetcode.com/problems/contains-duplicate/
// complexity: O[n]
func containsDuplicate(nums []int) bool {
	elementMap := map[int]bool{}
	for _, n := range nums {
		if _, exist := elementMap[n]; exist {
			return true
		} else {
			elementMap[n] = true
		}
	}
	return false
}
