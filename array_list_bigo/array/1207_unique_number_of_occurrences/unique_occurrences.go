package unique_number_of_occurrences

/*
	https://leetcode.com/problems/unique-number-of-occurrences/description/

*/

func uniqueOccurrences(arr []int) bool {
	freqMap := make(map[int]int)

	for _, num := range arr {
		freqMap[num]++
	}

	occurrenceMap := make(map[int]int)
	for _, v := range freqMap {
		occurrenceMap[v]++
	}

	if len(freqMap) == len(occurrenceMap) {
		return true
	}

	return false
}
