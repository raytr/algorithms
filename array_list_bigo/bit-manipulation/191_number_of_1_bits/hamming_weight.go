package number_of_1_bits

//https://leetcode.com/problems/number-of-1-bits/
// Complexity: O(n)

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num&1 == uint32(1) {
			count++
		}
		num = num >> 1
	}
	return count
}
