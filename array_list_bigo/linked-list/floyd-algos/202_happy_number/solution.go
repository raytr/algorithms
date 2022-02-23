package happy_number

/*
	problem: https://leetcode.com/problems/happy-number

   use floy to solve this projct, if a number isn't happly number
   slow == fast

   while fast == 1 || slow == 1
       fast = fast.getNext(getNext(fast))
       slow = slow.getNext(slow)
       if fast == slow => this number can't change anymore

   return val.getNext == 1
*/
func isHappy(n int) bool {
	fast, slow := getNext(getNext(n)), getNext(n)
	for fast != 1 && fast != slow {
		fast = getNext(getNext(fast))
		slow = getNext(slow)
	}
	return getNext(fast) == 1
}

func getNext(val int) int {
	sum := 0
	for val > 9 {
		digi := val % 10
		sum += digi * digi
		val /= 10
	}
	sum += val * val
	return sum
}
