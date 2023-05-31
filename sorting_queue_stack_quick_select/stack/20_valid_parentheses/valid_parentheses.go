package valid_parentheses

//https://leetcode.com/problems/valid-parentheses/

// if stack is empty => put c to stack
// if not, check the last element, if {, then next character should is }

func isValid(s string) bool {
	stack := make([]string, 0, len(s))
	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, string(c))
		} else if string(stack[len(stack)-1]) == "{" && string(c) == "}" {
			stack = stack[0 : len(stack)-1]
		} else if string(stack[len(stack)-1]) == "[" && string(c) == "]" {
			stack = stack[0 : len(stack)-1]
		} else if string(stack[len(stack)-1]) == "(" && string(c) == ")" {
			stack = stack[0 : len(stack)-1]
		} else {
			stack = append(stack, string(c))
		}
	}

	//if stack is empty
	if len(stack) == 0 {
		return true
	}
	return false
}
