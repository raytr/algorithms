package letter_combinations_of_a_phone_number

/*
	https://leetcode.com/problems/letter-combinations-of-a-phone-number/
	N is length of digits
	time complexity is O(4^N)

*/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	queue := []string{""}
	dic := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	for _, digit := range digits { // O(N)
		size := len(queue)
		if size > 0 {
			//with each element in queue
			for i := 0; i < size; i++ {
				//pop each element to make a path
				path := queue[0]
				queue = queue[1:]

				for _, char := range dic[digit] { //O(4) because the worst case is a number present for 4 characters
					queue = append(queue, path+string(char))
				}
			}
		}
	}

	return queue
}
