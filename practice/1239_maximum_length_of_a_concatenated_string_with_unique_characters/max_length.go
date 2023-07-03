package maximum_length_of_a_concatenated_string_with_unique_characters

/*
	https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/

	i check which concatenation has duplicate characters then remove this concatenation in arr

	sort it follow length of concatenation

	take a concatenation and gather it with each the rest concatenation in arr,
	- if exist duplicate character -> reject the result and move to next concatenation
	- otherwise, gather 2 concatenation and keep moving

	M is the length of concatenation that longgest
	N is the number of array

	K = M * N

	time complexity: O(2^k)
*/

func maxLength(arr []string) int {
	longestLen := 0
	result := []string{""}

	for _, word := range arr {
		resultLen := len(result)
		for i := 0; i < resultLen; i++ {
			newRes := result[i] + word
			if !isValid(newRes) {
				continue
			}

			//add to result
			result = append(result, newRes)
			longestLen = getMax(longestLen, len(newRes))
		}
	}

	return longestLen
}

func isValid(s string) bool {
	m := make(map[int32]int)
	for _, c := range s {
		m[c]++
		if m[c] > 1 {
			return false
		}
	}
	return true
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
