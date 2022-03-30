package minimum_number_of_flips_to_make_the_binary_string_alternating_ipt

/*
problem: https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/

clarify:
	is this string include 1 and 0?
	what is the maximum of length of this string?

we have 2 ways to make this string becomes internating
1: flip first character and increate count1 -> call doBecomeAlternating
2: call doBecomeAlternating



func doBecomeAlternating and return count

    run over this string, if prev == cur,
        while [0] == [1] && [0] != [last]
            do remove
        flip cur and increate count
        return count

*/

func minFlips(s string) int {
	r1 := []rune(s)
	r2 := []rune(s)

	count1 := doBecomeAlternating(r1, 0)

	r2[0] = flip(r2[0])
	count2 := doBecomeAlternating(r2, 0)

	return getMin(count1, count2)
}

func doBecomeAlternating(r []rune, count int) int {
	n := len(r)
	for i := 1; i < len(r); i++ {
		for r[0] == r[1] && r[0] != r[n-1] {
			r[0], r[n-1] = r[n-1], r[0]
		}

		//01001001101

		a, b := r[i], r[i-1]
		if a == b {
			c := flip(r[i])
			r[i] = c
			count++
		}
	}
	return count
}

func flip(c int32) int32 {
	if c == '1' {
		return '0'
	}
	return '1'
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
