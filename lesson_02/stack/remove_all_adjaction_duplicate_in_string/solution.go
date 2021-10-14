//package remove_all_adjaction_duplicate_in_string
package main

//problem: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/

func removeDuplicates(s string) string {
	stack := make([]int32, 0, len(s))
	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, c)
		} else if stack[len(stack)-1] == c {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}

	s = ""
	for _, c := range stack {
		s += string(c)
	}
	return s
}
