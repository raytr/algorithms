package rpn_calculator

import (
	"fmt"
	"math"
	"strconv"
)

func Calculate(expression string) (float64, error) {
	stack := make([]string, 0, len(expression))
	for _, c := range expression {
		if len(stack) < 2 {
			stack = append(stack, string(c))
		} else if string(c) == " " {
			continue
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
		} else if string(c) == "^" {
			num1, _ := strconv.Atoi(stack[len(stack)-2])
			num2, _ := strconv.Atoi(stack[len(stack)-1])
			result := math.Pow(float64(num1), float64(num2))
			stack = stack[0 : len(stack)-2]
			stack = append(stack, fmt.Sprintf("%v", result))
		} else {
			stack = append(stack, string(c))
		}
	}

	result, _ := strconv.ParseFloat(stack[0], 64)
	return result, nil

}
