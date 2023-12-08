package fibonacci_number

/*
	problem: https://leetcode.com/problems/fibonacci-number/
*/

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	sum, a, b := 0, 0, 1

	for i := 2; i <= n; i++ {
		sum = a + b
		a, b = b, sum
	}
	return sum
}
