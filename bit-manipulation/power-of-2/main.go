package power_of_2

//https://leetcode.com/problems/power-of-two/

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 1 {
		if remainder := n % 2; remainder == 1 {
			return false
		}
		n = n / 2
	}
	return true
}
