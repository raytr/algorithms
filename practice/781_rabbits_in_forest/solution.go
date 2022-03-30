package rabbits_in_forest

func numRabbits(answers []int) int {
	/*
	   we count how many rabbit answer the same answer
	   with each answer, we well max(number rabbit, answer)
	*/
	count := 0
	sameAnswerCount := make(map[int]int)

	for _, answer := range answers {
		sameAnswerCount[answer+1]++ //+1 for the rabbit is answering
		if answer+1 == sameAnswerCount[answer+1] {
			count += answer + 1
			delete(sameAnswerCount, answer+1)
		}
	}

	for k, v := range sameAnswerCount {
		count += Max(k, v)
	}
	return count
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
