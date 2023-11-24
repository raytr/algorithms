package evaluate_reverse_polish_notation

import (
	"fmt"
	"strconv"
)

/*
	problem: https://leetcode.com/problems/evaluate-reverse-polish-notation/

	create stack
	 if meet +|-|*|/, take number and operate them, and drop those 2
	 [2,1,+] => [2+1==3] => [3]

	time complexity: O(N), N is the length of tokens
*/

func evalRPN(tokens []string) int {
	stack := make([]string, 0, len(tokens))
	for _, c := range tokens {
		if len(stack) < 2 {
			stack = append(stack, c)
		} else if string(c) == "+" {
			num1, _ := strconv.Atoi(stack[len(stack)-2])
			num2, _ := strconv.Atoi(stack[len(stack)-1])
			result := num1 + num2
			stack = stack[0 : len(stack)-2]
			stack = append(stack, fmt.Sprintf("%v", result))
		} else if string(c) == "-" {
			num1, _ := strconv.Atoi(stack[len(stack)-2])
			num2, _ := strconv.Atoi(stack[len(stack)-1])
			result := num1 - num2
			stack = stack[0 : len(stack)-2]
			stack = append(stack, fmt.Sprintf("%v", result))
		} else if string(c) == "*" {
			num1, _ := strconv.Atoi(stack[len(stack)-2])
			num2, _ := strconv.Atoi(stack[len(stack)-1])
			result := num1 * num2
			stack = stack[0 : len(stack)-2]
			stack = append(stack, fmt.Sprintf("%v", result))
		} else if string(c) == "/" {
			num1, _ := strconv.Atoi(stack[len(stack)-2])
			num2, _ := strconv.Atoi(stack[len(stack)-1])
			result := num1 / num2
			stack = stack[0 : len(stack)-2]
			stack = append(stack, fmt.Sprintf("%v", result))
		} else {
			stack = append(stack, c)
		}
	}

	result, _ := strconv.Atoi(stack[0])
	return result
}
