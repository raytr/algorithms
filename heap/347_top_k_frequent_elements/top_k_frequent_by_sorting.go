package top_k_frequent_elements

import "sort"

/*
	create freqMap with map[num]frequent  (O(n))
	integrate all numbers to fulfill the map
    create []freqVal and push all to this
	sort desc this array (O(N * logN))
	pop k  val to push to result

	time complexity: O(N * logN)
*/

func topKFrequentBySorting(nums []int, k int) []int {

	m := make(map[int]int)
	for _, n := range nums { //O(n)
		m[n]++
	}

	freqVals := make([]FreqVal, 0, len(m))
	for num, freq := range m {
		freqVals = append(freqVals, FreqVal{
			val:  num,
			freq: freq,
		})
	}

	//sort desc: O(n * logN)
	sort.Slice(freqVals, func(i, j int) bool {
		return freqVals[i].freq > freqVals[j].freq
	})

	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, freqVals[i].val)
	}

	return result
}
