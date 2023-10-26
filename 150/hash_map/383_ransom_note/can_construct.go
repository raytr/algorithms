package _83_ransom_note

/*
	problem: https://leetcode.com/problems/ransom-note

	ransomNoteMap: map[character]int
	magazineMap: map[character]int

	ransomNoteMap[a] = 2

	magazineMap[a] = 2
	magazineMap[b] = 1

	time complexity: O(N). N is length of magazine
*/

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	ransomNoteMap := make(map[rune]int)
	magazineMap := make(map[rune]int)

	for _, c := range magazine {
		magazineMap[c]++
	}

	for _, c := range ransomNote {
		ransomNoteMap[c]++
	}

	for k, _ := range ransomNoteMap {
		if magazineMap[k] < ransomNoteMap[k] {
			return false
		}
	}

	return true
}
