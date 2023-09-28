package valid_palindrome

import (
	"strings"
)

/*
	problem: https://leetcode.com/problems/valid-palindrome/
*/

func isPalindrome(s string) bool {
	/*
	   dic -> 24leters
	   result (string),


	   left, right ->


	   a  b. b. a
	      l  r
	*/

	dic := newAlphabetDic()
	ns := make([]int32, 0)
	s = strings.ToLower(s)
	for _, c := range s {
		if _, exist := dic[c]; exist {
			ns = append(ns, c)
		}
	}

	left, right := 0, len(ns)-1
	for left < right {
		if ns[left] != ns[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func newAlphabetDic() map[int32]int {
	dic := make(map[int32]int)
	dic['a']++
	dic['b']++
	dic['c']++
	dic['d']++
	dic['e']++
	dic['f']++
	dic['g']++
	dic['h']++
	dic['i']++
	dic['j']++
	dic['k']++
	dic['l']++
	dic['m']++
	dic['n']++
	dic['o']++
	dic['p']++
	dic['q']++
	dic['r']++
	dic['s']++
	dic['t']++
	dic['u']++
	dic['v']++
	dic['w']++
	dic['x']++
	dic['y']++
	dic['z']++
	dic['0']++
	dic['1']++
	dic['2']++
	dic['3']++
	dic['4']++
	dic['5']++
	dic['6']++
	dic['7']++
	dic['8']++
	dic['9']++
	return dic
}
