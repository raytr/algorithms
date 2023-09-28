package power_of_three

/*
problem: https://leetcode.com/problems/power-of-three/description/
what is minimum and maximum of n
time complexity: O(log 3 ofN)
*/
func isPowerOfThree(n int) bool {

	if n < 1 {
		return false
	}

	powerOfThree := 1

	for powerOfThree <= n {
		if powerOfThree == n {
			return true
		}
		powerOfThree *= 3
	}

	return false
}
