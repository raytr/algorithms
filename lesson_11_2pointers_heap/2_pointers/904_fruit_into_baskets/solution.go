package fruit_into_baskets

/*
		problem: https://leetcode.com/problems/fruit-into-baskets


	   0 1 2 2
	   s f     => 2 types (it enough)
	     s f   => 3 types => delete(fruits[s]), s++
	     while fruits[]


	   2 pointers f,s
	   f move from head of fruits to end of fruits
	   if freqMap[fruits[f]] is exist || len(freqMap) <= 2
	       f++
	   otherwise
	       decrease freqMap
	       s++

	   complexity: O(n)
*/

func totalFruit(fruits []int) int {
	max, s, f := 1, 0, 0
	freqMap := make(map[int]int)

	for f < len(fruits) {
		if _, exist := freqMap[fruits[f]]; exist || len(freqMap) <= 1 {
			freqMap[fruits[f]]++
			max = Max(max, f-s+1)
			f++
		} else {
			descreaseMap(freqMap, fruits[s])
			s++
		}
	}
	return max
}

func descreaseMap(m map[int]int, key int) {
	if m[key] == 1 {
		delete(m, key)
	} else {
		m[key]--
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
