package fruit_into_baskets

/*
	problem: https://leetcode.com/problems/fruit-into-baskets


	solution:
	2 pointers: fast and slow
	fast move from head of fruits to end of fruits
	if freqMap[fruits[fast]] is exist || len(freqMap) <= 1
		add fruits[fast] to freqMap
	   f++
	otherwise
	   decrease freqMap
	   s++

	complexity: O(n)

	s
	0   1   2   1   2   0   1   2
	f

	max = 4
*/

func totalFruit(fruits []int) int {
	slow, fast, max := 0, 0, 1
	fruitFreqMap := make(map[int]int)

	for fast < len(fruits) {
		if _, exist := fruitFreqMap[fruits[fast]]; exist || len(fruitFreqMap) <= 1 {
			fruitFreqMap[fruits[fast]]++
			currentNumberOfFruits := fast - slow + 1
			max = getMax(max, currentNumberOfFruits)
			fast++
		} else {
			decreaseMap(fruitFreqMap, fruits[slow])
			slow++
		}
	}

	return max
}

func decreaseMap(m map[int]int, key int) {
	if m[key] == 1 {
		delete(m, key)
	} else {
		m[key]--
	}
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
