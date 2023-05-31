package main

import "fmt"

/*
	the following Factorial function writen by your colleague Fred is supposed to return the factorial of a number, but it is broken.
	fix the factorial function

	as a reminder: factorial(n) = 1 * 2 * 3 * ... * n
*/

func main() {
	fmt.Println(Factorial(6))
}

func Factorial(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return n * Factorial(n-1)
}
