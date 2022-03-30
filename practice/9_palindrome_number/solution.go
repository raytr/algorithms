package palindrome_number

import "fmt"

/*

	problem: https://leetcode.com/problems/palindrome-number/

   1 2 3 2 1
   l.      r
     l.  r
       lr

   l.Val == r.Val => l++, r--

   l == r => return true
   if l != r && l.Val != r.Val => return false

   run time : O(n)
   space complexity: O(1)

   proun: check,
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x > 0 && x < 10 {
		return true
	}

	str := fmt.Sprintf("%v", x)
	l, r := 0, len(str)-1 //l ==0, ; r == 4

	for l <= r {
		if str[l] == str[r] { //l.val == 1, r.val 1
			l++
			r--
		} else {
			return false
		}
	}

	if l != r && str[l] != str[r] { //l == 2, r == 2; l.Val == 3; r.Val == 3
		return false
	}

	return true
}
