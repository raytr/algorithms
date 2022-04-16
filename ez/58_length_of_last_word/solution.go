package length_of_last_word

/*
	link: https://leetcode.com/problems/length-of-last-word/

   clarify:
       this string just has alb and space
       how about " "

   iterate from right, with each iterate, we increate count
   if we meet " ", if count > 0, return count, other wise => do nothing

	O(n)
*/

func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			count++
		} else if count > 0 {
			return count
		}
	}
	return count
}
