package single_number_ii

/*
	https://leetcode.com/problems/single-number-ii/

*/

func singleNumber(nums []int) int {
	freqMap := make(map[int]int)
	result := 0

	for _, num := range nums {
		freqMap[num]++
		if freqMap[num] == 3 {
			delete(freqMap, num)
		}
	}

	for k, _ := range freqMap {
		result = k
	}

	return result
}
