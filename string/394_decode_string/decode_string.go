package decode_string

import (
	"strconv"
)

/*
https://leetcode.com/problems/decode-string/

stack = 3 [ a
go though all character

	if c == ]
		pop character out of stack util meet [ and push all characters to decodedString
		pop [
		pop the next character (k), that is digit character util meet the letter
		decodedString * k times
		push all decodedString to the stack

read the stack and connect to the result (string)

	because go thought each character 1 time -> runtime is O(n)
	we need memory to store the stack, decodedString -> space is O(2n) -> O(n)
*/

func decodeString(s string) string {
	stack := make([]int32, 0, len(s))
	decodedString := ""
	isValid := false
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			isValid = true
		}

		if len(stack) > 0 && c == ']' {
			for stack[len(stack)-1] != '[' {
				//pop the last
				theLast := string(stack[len(stack)-1])
				stack = stack[0 : len(stack)-1]
				decodedString = theLast + decodedString
			}

			//pop '['
			stack = stack[0 : len(stack)-1]

			//pop the digit
			num := ""
			for len(stack) > 0 && isNum(string(stack[len(stack)-1])) {
				k := string(stack[len(stack)-1])
				stack = stack[0 : len(stack)-1]
				num = k + num
			}
			decodedString = mulString(decodedString, num)

			//push all decodedString to stack
			for _, decodedChar := range decodedString {
				stack = append(stack, decodedChar)
			}

			decodedString = ""
		} else {
			stack = append(stack, c)
		}
	}

	if !isValid {
		return ""
	}
	result := ""
	for _, c := range stack {
		result += string(c)
	}

	return result
}

func mulString(s, timeStr string) string {
	time, _ := strconv.Atoi(timeStr)
	result := s
	for i := 1; i < time; i++ {
		result += s
	}
	return result
}

func isNum(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}
