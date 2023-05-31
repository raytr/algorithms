package basic_calculator_ii

import (
	"fmt"
	"strconv"
)

// can s empty
// how about length of string
// what is your expectation about complexity
// what result if 44+1 =>

//
// s1: remove all space
// s2: convert string to array that include number and sign : " 42 +1 " => [42,+,1]
// we will separate 2 group is [*, /] and [+, =]
// loop s 2 time, 1 for [*, /] , 1 for [+, -]
// if c != * | / | + | - => push => stack
// if c == * | / => result = c[i-1] *|/ c[i+1]
// similarity for + | -
// after all we have result
// complexity : O(2 length of s)

//calculate("3+2*2")
//calculate(" 3/2 ")
//calculate(" 3+5 / 2 ")
//calculate("42")
//calculate("0")
//calculate("1+1+1")

func calculate(s string) int {

	arr := make([]string, 0, len(s))
	temp := ""

	// remove all space and push
	for _, c := range s {
		if string(c) == " " {
			continue
		}
		arr = append(arr, temp+string(c))
	}

	stack := make([]string, 0, len(arr))

	// convert string to array that include number and sign
	for i, c := range arr {
		if string(c) == "+" || string(c) == "-" || string(c) == "*" || string(c) == "/" {
			if temp != "" {
				stack = append(stack, temp)
				temp = ""
			}
			stack = append(stack, c)
		} else if i == len(arr)-1 { // the last character
			stack = append(stack, temp+c)
		} else {
			temp += c
		}
	}

	arr = stack
	stack = make([]string, 0, len(arr))

	//first loop for *|/
	for i := 0; i < len(arr); i++ {
		if len(stack) == 0 {
			stack = append(stack, arr[i])
		} else if arr[i] == "*" {
			num1, _ := strconv.Atoi(stack[len(stack)-1])
			num2, _ := strconv.Atoi(arr[i+1])
			result := num1 * num2
			stack = stack[0 : len(stack)-1]
			stack = append(stack, fmt.Sprintf("%v", result))
			i++ //because i used the next element
		} else if arr[i] == "/" {
			num1, _ := strconv.Atoi(stack[len(stack)-1])
			num2, _ := strconv.Atoi(arr[i+1])
			result := num1 / num2
			stack = stack[0 : len(stack)-1]
			stack = append(stack, fmt.Sprintf("%v", result))
			i++ //because i used the next element
		} else {
			stack = append(stack, arr[i])
		}
	}

	// for + | -
	arr = stack
	stack = []string{}
	isCal := false
	for i := 0; i < len(arr); i++ {
		if len(stack) == 0 {
			stack = append(stack, arr[i])
		} else if arr[i] == "+" {
			num1, _ := strconv.Atoi(stack[len(stack)-1])
			num2, _ := strconv.Atoi(arr[i+1])
			result := num1 + num2
			stack = stack[0 : len(stack)-1]
			stack = append(stack, fmt.Sprintf("%v", result))
			i++
			isCal = true
		} else if arr[i] == "-" {
			num1, _ := strconv.Atoi(stack[len(stack)-1])
			num2, _ := strconv.Atoi(arr[i+1])
			result := num1 - num2
			stack = arr[0 : len(stack)-1]
			stack = append(stack, fmt.Sprintf("%v", result))
			i++
			isCal = true
		} else {
			stack = append(stack, arr[i])
		}
	}

	var num int
	if isCal {
		num, _ = strconv.Atoi(stack[0])
	} else {
		num, _ = strconv.Atoi(arr[0])
	}
	return num
}
