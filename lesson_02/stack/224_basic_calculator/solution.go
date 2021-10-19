//package basic_calculator
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calculate("1 + 1"))
	fmt.Println(calculate(" 2-1 + 2 "))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}

//problem: leetcode.com/problems/basic-calculator/?fbclid=IwAR1vwWv1xF8E3CyWuOZE12O5V-cnQKdxWn9qSyYfmBz1m6DEhHaZC81r9oA

/*
	use loop to remove space and get number and operator. Ex "  1+1 " => ["1","+","1"]; "43+1" => ["43","+","1"]
	define a function calculateIn to calculate: ["1","+","1"] = 2

	loop from lefrt to right over s. If meet ( => create a new stack to store number and overrator, see ) will finish the stack
	=> 1 - (1+2) => newStack have ["1","+","2"] and call calculateIn function to get result => remove pair of ()
	in the end, we will has a stack has only 1 element => result
*/

func calculate(s string) int {
	arr := make([]string, 0, len(s))
	tempNum := ""
	for _, c := range s {
		if string(c) == " " || string(c) == "(" || string(c) == ")" {
			if tempNum != "" {
				arr = append(arr, tempNum)
				tempNum = ""
			}
			continue
		} else if string(c) == "+" || string(c) == "-" {
			if tempNum != "" {
				arr = append(arr, tempNum)
				tempNum = ""
			}
			arr = append(arr, string(c))
		} else {
			tempNum += string(c)
		}
	}
	if tempNum != "" {
		arr = append(arr, tempNum)
		tempNum = ""
	}

	var stack []string
	for i := 0; i < len(arr); i++ {
		lastElementInStack := len(stack) - 1
		if arr[i] == "+" {
			num1, _ := strconv.Atoi(stack[lastElementInStack])
			num2, _ := strconv.Atoi(arr[i+1])
			sum := num1 + num2
			stack = stack[:lastElementInStack]
			stack = append(stack, strconv.Itoa(sum))
			i++
		} else if arr[i] == "-" {
			num1, _ := strconv.Atoi(stack[lastElementInStack])
			num2, _ := strconv.Atoi(arr[i+1])
			sum := num1 - num2
			stack = stack[:lastElementInStack]
			stack = append(stack, strconv.Itoa(sum))
			i++
		} else {
			stack = append(stack, arr[i])
		}
	}
	result, _ := strconv.Atoi(stack[0])
	return result
}
