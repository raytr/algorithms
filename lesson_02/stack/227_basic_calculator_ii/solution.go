//package basic_calculator_ii
package main

import (
	"fmt"
	"strconv"
)

// can s empty
// how about length of string
// what is your expectation about complexity
// create stack
// we will separate 2 group is [*, /] and [+, =]
// loop s 2 time, 1 for [*, /] , 1 for [+, -]
// if c != * | / | + | - => push => stack
// if c == * | / => result = c[i-1] *|/ c[i+1]
// similarity for + | -
// after all we have result
// complexity : O(2 length of s)

func main() {
	//fmt.Println(calculate("3+2*2"))
	//fmt.Println(calculate(" 3/2 "))
	//fmt.Println(calculate(" 3+5 / 2 "))
	fmt.Println(calculate("42"))
	//fmt.Println(calculate("0"))
}

func calculate(s string) int {

	arr := make([]string, 0, len(s))
	temp := ""

	for i, c := range s {
		if string(c) == " " {
			continue
		}
		if string(c) == "+" || string(c) == "-" || string(c) == "*" || string(c) == "/" {
			arr = append(arr, temp)
			temp = ""
			arr = append(arr, string(c))
		} else if i == len(s)-1 {
			arr = append(arr, temp+string(c))
		} else {
			temp += string(c)
		}
	}

	stack := make([]string, 0, len(s))
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
			stack = stack[0 : i-1]
			stack = append(stack, fmt.Sprintf("%v", result))
			i++
			isCal = true
		} else if arr[i] == "-" {
			num1, _ := strconv.Atoi(stack[len(stack)-1])
			num2, _ := strconv.Atoi(arr[i+1])
			result := num1 - num2
			stack = arr[0 : i-1]
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
