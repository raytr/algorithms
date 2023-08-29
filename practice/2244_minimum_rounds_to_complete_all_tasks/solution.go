package minimum_rounds_to_complete_all_tasks

/*
	problem: https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/

	tasks = [2,2,3,3,2,4,4,4,4,4]

	numFreq = map[num]freq => O(n)
	numFreq[2] = 3
	numFreq[3] = 2
	numFreq[4] = 5

	3 case
	if freq % 3 == 0 => res = req / 3
	else if frq % 2 == 0 => req / 2
	else
		remainder = freq % 3
		res = freq / 3
		if remainder % 2 == 0 {
			res = res + (remainder / 2)
		else return -1

	return count

	time complexity: O(n)
*/

func minimumRounds(tasks []int) int {
	count := 0
	numFreq := make(map[int]int)

	for _, num := range tasks {
		numFreq[num]++
	}

	for _, numLen := range numFreq {
		if numLen%3 == 0 {
			count += numLen / 3
		} else if numLen%2 == 0 {
			count += numLen / 2
		} else {
			remainder := numLen % 3
			count += numLen / 3
			if remainder%2 == 0 {
				count += remainder / 2
			} else {
				return -1
			}
		}
	}

	return count
}
