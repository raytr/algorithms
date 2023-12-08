package edit_distance

func minDistanceRecursion(word1 string, word2 string) int {
	return minDistanceRecur(word1, word2, len(word1), len(word2))
}

func minDistanceRecur(word1 string, word2 string, word1Index int, word2Index int) int {
	if word1Index == 0 {
		return word2Index
	}
	if word2Index == 0 {
		return word1Index
	}
	if word1[word1Index-1] == word2[word2Index-1] {
		return minDistanceRecur(word1, word2, word1Index-1, word2Index-1)
	} else {
		insertOperation := minDistanceRecur(word1, word2, word1Index, word2Index-1)
		deleteOperation := minDistanceRecur(word1, word2, word1Index-1, word2Index)
		replaceOperation := minDistanceRecur(word1, word2, word1Index-1, word2Index-1)
		return getMin(insertOperation, deleteOperation, replaceOperation) + 1
	}
}

func getMin(a, b, c int) int {
	min := a
	if b < a {
		min = b
	}
	if c < b {
		min = c
	}
	return min
}
