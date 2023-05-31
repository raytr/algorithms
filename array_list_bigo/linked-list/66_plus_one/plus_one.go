package plus_one

/*
	https://leetcode.com/problems/plus-one/description/
	move from tail to head -> if remember is true -> +1 for next num
	put to array

	run time : linear time
	space complexity: constant
*/

func plusOne(digits []int) []int {
	n := len(digits)
	remember := true // because +1 to the num

	for i := n - 1; i >= 0; i-- {
		if remember {
			num := digits[i] + 1
			if num == 10 {
				digits[i] = 0
				remember = true
			} else {
				digits[i] = num
				remember = false
			}
		}
	}

	if remember {
		digits = append([]int{1}, digits...)
	}

	return digits
}
