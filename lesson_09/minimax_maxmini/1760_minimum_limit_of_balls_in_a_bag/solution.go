package minimum_limit_of_balls_in_a_bag

/*

problem: https://leetcode.com/problems/minimum-limit-of-balls-in-a-bag/
because this is maxmini problem so :

fist of all: find the max
hi = max in nums
lo = min in nums

THE SMALLER THE MID, THE HIGHER THE OPES
THE HIGHER THE MID, THE SMALLER THE OPES

bin: mid = min + (max-min)/2

opes := 0

while score min <= max
	interate all of numbers of ball in bag
	if num % score != 0 => opes += num/score
	if num % score == 0 => opes += num/score - 1

	if ops <= maxOpes (no need care the cases has < ops => find smaller mid) {
		res = mid
		lo = mid + 1

	} else {
		hi = mid - 1
	}

return mid



*/

func minimumSize(nums []int, maxOperations int) int {
	lo := 0
	hi := 0
	var res int

	for _, num := range nums {
		hi += num
	}

	for lo <= hi {
		opes := 0
		mid := lo + (hi-lo)/2
		for _, num := range nums {
			if num%mid != 0 {
				opes += num / mid
			} else {
				opes += num/mid - 1
			}
			if opes > maxOperations {
				lo = mid + 1
				break
			}
		}
		if opes <= maxOperations {
			res = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return res
}
