package first_unique_character_in_a_string

//problem: https://leetcode.com/problems/first-unique-character-in-a-string/?fbclid=IwAR2sAS2bEdWs70pZil4i9FcYO-1NrUPvOF4zM39JDV9PzklnPbgVoqudI0k

/*
	create map with key is frequency, value
	loop over the s, count when meet each character
	loop map => if value != 1 => delete it
	loop over the s, each iterate, check this character exist in charMap => if yes => return index | no => continue

	complexity: O(length of s}
*/
func firstUniqChar(s string) int {
	charMap := make(map[string]int)
	for _, c := range s {
		charMap[string(c)]++
	}
	for k, v := range charMap {
		if v != 1 {
			delete(charMap, k)
		}
	}
	for i, c := range s {
		if _, exist := charMap[string(c)]; exist {
			return i
		}
	}
	return -1
}
