package main

import (
	"fmt"
)

func printSubsequence(output, input string) {

	//# Base Case
	//# if the input is empty print the output string
	if len(input) == 0 {
		fmt.Println(output)
		return
	}

	//# output is passed with including the
	//# 1st character of input string
	printSubsequence(input[1:], output+string(input[0]))

	//# output is passed without including the
	//# 1st character of input string
	printSubsequence(input[1:], output)
}

func main() {
	strData := "XYZ"
	printSubsequence("", strData)
	fmt.Println()
}
