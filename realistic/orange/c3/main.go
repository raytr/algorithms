package main

import "strings"

func main() {
	Solution()
}

func Solution(T []string) int {
	freq := make(map[string][]int)
	for i, s := range T {
		r1 := []rune(s)
		r1[0], r1[1] = r1[1], r1[0]
		freq[string(r1)] = append(freq[string(r1)], i)

		r2 := []rune(s)
		r2[0], r2[1] = r2[1], r2[0]
		freq[string(r1)] = append(freq[string(r1)], i)

	}
}

func newString(s string) string {
	var sb strings.Builder
	if _, err := sb.WriteString(s); err != nil {
		//log & return
	}
	return sb.String()
}
