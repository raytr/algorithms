package n_th_tribonaccia_number

/*
	problem: https://leetcode.com/problems/n-th-tribonacci-number/
*/

func tribonacci(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	sum, a, b, c := 0, 0, 1, 1
	for i := 4; i < n+1; i++ {
		a, b, c = b, c, a+b+c
		sum = a + b + c
	}
	return sum
}
