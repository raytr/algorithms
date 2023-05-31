package RPNcalculator

import (
	"errors"
	"strconv"
	"strings"
)

func Calculate(expression string) (float64, error) {
	//convert to an array of num and operation
	arr := strings.Split(expression, " ")
	stack := make([]float64, 0, len(arr))

	for _, str := range arr {
		if len(stack) < 2 {
			num, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return 0, errors.New("wrong format")
			}
			stack = append(stack, num)
		} else if str == "*" {
			//pop from stack
			n := len(stack)
			num1 := stack[n-2]
			num2 := stack[n-1]
			stack = stack[:n-2]
			res := num1 * num2
			stack = append(stack, res)
		} else if str == "/" {
			n := len(stack)
			num1 := stack[n-2]
			num2 := stack[n-1]
			stack = stack[:n-2]
			res := num1 / num2
			stack = append(stack, res)
		} else if str == "+" {
			n := len(stack)
			num1 := stack[n-2]
			num2 := stack[n-1]
			stack = stack[:n-2]
			res := num1 + num2
			stack = append(stack, res)
		} else if str == "-" {
			n := len(stack)
			num1 := stack[n-2]
			num2 := stack[n-1]
			stack = stack[:n-2]
			res := num1 - num2
			stack = append(stack, res)
		} else {
			num, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return 0, errors.New("wrong format")
			}
			stack = append(stack, num)
		}
	}
	return stack[0], nil
}
