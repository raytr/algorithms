package fibonacci_number

/*
	problem: https://leetcode.com/problems/fibonacci-number/
*/

// bottom up
func fibBottomUp(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	res, a, b := 0, 0, 1
	for i := 2; i <= n; i++ {
		res = a + b
		a = b
		b = res
	}
	return res
}

// top down
var fibo map[int]int

func fibTopDown(n int) int {
	fibo = make(map[int]int)
	for i := 0; i <= n; i++ {
		fibo[i] = -1
	}

	return cal(n)
}

func cal(i int) int {
	// base case
	if i == 1 || i == 2 {
		return 1
	}

	if fibo[i] != -1 {
		return fibo[i]
	}
	fibo[i] = cal(i-1) + cal(i-2)
	return fibo[i]
}
