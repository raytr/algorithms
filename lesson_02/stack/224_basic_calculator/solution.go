//package basic_calculator
package main

import "fmt"

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
		if string(c) == " " {
			continue
		}
		if string(c) == "(" || string(c) == ")" || string(c) == "+" || string(c) == "-" {
			arr = append(arr, tempNum)
			arr = append(arr, string(c))
			tempNum = ""
		} else {
			tempNum += string(c)
		}
	}

}

func calculateBrackets(s []string) int {
	for _, c := range s {
		stackInBrackets := []string{}
		if c == "(" {

		}
	}
}
