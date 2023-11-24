package two_sum

//problem: https://leetcode.com/problems/two-sum/

/*
	q: is array has negative number

	theRestMap
	iterate nums
		check theRestMap[target - num] is exist ;
			if not => theRestMap[target - num] = index
			if exist => [theRestMap[target - num], i]

*/
// complexity: O(n)
func twoSum(nums []int, target int) []int {
	theRestMap := make(map[int]int)
	for i, num := range nums {
		if val, exist := theRestMap[num]; exist {
			return []int{val, i}
		} else {
			theRestMap[target-num] = i
		}
	}
	return []int{}
}
