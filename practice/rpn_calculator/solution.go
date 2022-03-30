package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

/*

	Implement the RPN calculator.Calculate function, which evaluates expressions using Reverse Polish Notation.
	The function takes as input a string of RPN expressions along with float64 numbers as tokens, separated by spaces.
	The function must return the result of the evaluation as a float64 or an error indicating invalid input.

	An RPN Expression may contain the following operators.

	+, - (addition, subtraction)
	*, / (multiplication, division)
	^ (exponentiation)
	not (negation, not(x) should return -x)
	in (in(a, b, c) should return 1 if b >= a >= c, otherwise 0)

	Examples
		Calculate("-62 -52 ^ 8 -41 in") should return 1
		Calculate("42") should return 42
		Calculate("-19 not * -68 + not") should return an error

*/

func Calculate(expression string) (float64, error) {
	expression += " "
	stack := make([]float64, 0)
	tempStr := ""

	for i := 0; i < len(expression); i++ {
		c := expression[i]
		a := string(c)
		fmt.Println(a)
		if string(c) == " " {
			if tempStr != "" {
				num, _ := strconv.ParseFloat(tempStr, 64)
				stack = append(stack, num)
				tempStr = ""
			}
			continue
		} else if len(stack) < 2 && tempStr == "" {
			tempStr += string(c)
		} else if string(c) == "+" && string(expression[i-1]) == " " && string(expression[i+1]) == " " {
			if len(stack) > 2 {
				return 0, errors.New("invalid")
			}
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			result := num1 + num2
			stack = stack[0 : len(stack)-2]
			stack = append(stack, result)
		} else if string(c) == "-" && string(expression[i-1]) == " " && string(expression[i+1]) == " " {
			if len(stack) > 2 {
				return 0, errors.New("invalid")
			}
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			result := num1 - num2
			stack = stack[0 : len(stack)-2]
			stack = append(stack, result)
		} else if string(c) == "*" && string(expression[i-1]) == " " && string(expression[i+1]) == " " {
			if len(stack) > 2 {
				return 0, errors.New("invalid")
			}
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			result := num1 * num2
			stack = stack[0 : len(stack)-2]
			stack = append(stack, result)
		} else if string(c) == "/" && string(expression[i-1]) == " " && string(expression[i+1]) == " " {
			if len(stack) > 2 {
				return 0, errors.New("invalid")
			}
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			result := num1 / num2
			stack = stack[0 : len(stack)-2]
			stack = append(stack, result)
		} else if string(c) == "^" && string(expression[i-1]) == " " && string(expression[i+1]) == " " {
			if len(stack) > 2 {
				return 0, errors.New("invalid")
			}
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			result := math.Pow(float64(num1), float64(num2))
			stack = stack[0 : len(stack)-2]
			stack = append(stack, result)
		} else {
			tempStr += string(c)
		}
	}

	result := stack[0]
	return result, nil

}
