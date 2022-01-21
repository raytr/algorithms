package remove_all_adjaction_duplicate_in_string

/*
	problem: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/


   check each character (c) from s
   stack, if stack empty => push to stack
   if the last element == c => remove last element
   return stack

	time complexity: O(n) because traversal all character of s
	space complexity: O(n) create stack to store node

*/
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
