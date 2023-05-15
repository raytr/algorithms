package gray_code

/*
	https://leetcode.com/problems/gray-code/

	you can see in indicates imagy, the grey code sequence is
	first is 0, 1
	the next number follow the rule, num + 2 ,  num - 1
	for ex: 0, 1, 3, 2, 5, 4, 7, 6
*/

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	res := []int{0, 1}
	base := 2
	for i := 1; i < n; i++ {
		rlen := len(res)
		for j := rlen - 1; j >= 0; j-- {
			res = append(res, res[j]+base)
		}
		base = base * 2
	}
	return res

}
