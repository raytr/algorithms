package pelindrome_permutation

/*
	create a map with key is character and value is the time that character appear
	get all value in map mod to 2, if the just has only one character that has the remainder <= 1 => return true
*/

func canPermutePalindrome(s string) bool {
	charMap := make(map[string]uint)
	count := 0
	for _, c := range s {
		if _, exist := charMap[string(c)]; !exist {
			charMap[string(c)] = 1
		} else {
			charMap[string(c)] += 1
		}
	}

	for _, v := range charMap {
		if v%2 == 1 {
			count++
		}
	}

	if count <= 1 {
		return true
	}
	return false
}
