package basic_calculator

import (
	"strconv"
)

//problem: https://leetcode.com/problems/basic-calculator/

/*
	with this problem, we always want to calculate from of value of pair of inmost "()" to outmost "()"
	ex: (1+(2-1)) => first calculate the sum of pair of inmost () is: 2-1. After that: outmost "()" is 1+1
	so, we need to reverse to solve this problem.
	loop from right to left => remove space and check is digit and reformat it. Push all in to a stack
	if meet "(" => mean is finish inmost (). So, we need find ")" from right to left in stack and calculate it, after each
		iterate, remote the last element in stack

	the result will is stack[0]
*/

func calculate(s string) int {
	tempNum := ""
	stack := make([]string, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		if num, err := strconv.Atoi(string(s[i])); err == nil {
			if tempNum != "" {
				tempNum = strconv.Itoa(num) + tempNum
			} else {
				tempNum = strconv.Itoa(num)
			}
		} else {
			if tempNum != "" {
				stack = append(stack, tempNum)
				tempNum = ""
			}
			if string(s[i]) == " " {
				continue
			} else if string(s[i]) == "(" {
				res := 0
				res, stack = evalueteExpr(stack)
				stack = append([]string{}, stack[:len(stack)-1]...)
				stack = append(stack, strconv.Itoa(res))
			} else {
				stack = append(stack, string(s[i]))
			}
		}
	}
	if tempNum != "" {
		stack = append(stack, tempNum)
		tempNum = ""
	}

	result, _ := evalueteExpr(stack)
	return result
}

func evalueteExpr(stack []string) (int, []string) {
	temp := ""
	if stack[len(stack)-1] == "-" {
		temp = "-" + stack[len(stack)-2]
		stack = stack[:len(stack)-2]
	} else {
		temp = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	result, _ := strconv.Atoi(temp)
	for len(stack) > 0 && stack[len(stack)-1] != ")" {
		sign := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nextNum, _ := strconv.Atoi(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
		if sign == "+" {
			result += nextNum
		} else {
			result -= nextNum
		}
	}
	return result, stack
}
