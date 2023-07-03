package maximum_number_of_vowels_in_a_substring_of_given_length

/*
	https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
	integrate all character in s, if meet a vowels, get index
		check how many vowels in the substring that has length == 7 (if possiple) => update max
		continue check next character


	if max > k => return k
	max k < max => max
*/

func maxVowels(s string, k int) int {
	max := 0
	count := 0
	n := len(s)

	for i, c := range s {
		if c == 'u' || c == 'e' || c == 'a' || c == 'o' || c == 'i' {
			endIndex := i + k - 1
			if endIndex > n-1 {
				endIndex = n - 1
			}

			for curIdx := i; curIdx <= endIndex; curIdx++ {
				if s[curIdx] == 'u' || s[curIdx] == 'e' || s[curIdx] == 'a' || s[curIdx] == 'o' || s[curIdx] == 'i' {
					count++
				}
			}

			max = getMax(max, count)
			count = 0
		}
	}

	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
