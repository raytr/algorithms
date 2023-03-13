package main

import "strings"

func main() {
	Solution([]string{"aab", "cab", "baa", "baa"})
}

func Solution(T []string) int {
	// xuat ra tat ca kha nang
	freqM := make(map[string][]int)
	for i, s := range T {
		freqM[s] = append(freqM[s], i)

		r1 := []rune(s)
		r1[0], r1[1] = r1[1], r1[0]
		s1 := string(r1)
		freqM[s1] = append(freqM[s1], i)

		r2 := []rune(s)
		r2[1], r2[2] = r2[2], r2[1]
		s2 := string(r2)
		freqM[s2] = append(freqM[s2], i)
	}

	res := 0
	for _, v := range freqM {
		count := make(map[int]bool)
		for _, n := range v {
			count[n] = true
		}
		if len(count) > res {
			res = len(count)
		}
	}

	return res
}

func newString(s string) string {
	var sb strings.Builder
	if _, err := sb.WriteString(s); err != nil {
		//log & return
	}
	return sb.String()
}
