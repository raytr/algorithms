package longest_well_performing_interval

/*

	    revert array with 1 with hour > 8 and -1 for ortherwise

        iterate all of the hours to calculate prefixsum

		if last prefix sum > 0 => return all of array
		else {
			iterate all of prefixSums with while i
				if prxSum[i] > 0 => max++
				else if prxSum[i] == idxMap[prxSum[i]-1]
						=> m := i - prxSum[i] == idxMap[prxSum[i]-1]; find Max
						(for case that all of element prxSum subarray are negative number)
				else => idxMap[prxSum[i]-1] = i
		}

        return max

        time complexity: O(n)
		space complexity: O(1)
*/

func longestWPI(hours []int) int {
	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
	}

	//calculate prefix sum
	for i := 1; i < len(hours); i++ {
		hours[i] = hours[i-1] + hours[i]
	}

	max := 0
	idxMap := make(map[int]int)

	if hours[len(hours)-1] > 0 {
		return len(hours)
	}

	for i := 0; i < len(hours); i++ {
		if hours[i] == 1 {
			max = Max(max, i+1)
		} else if _, exist := idxMap[hours[i]]; exist {
			count := i - (idxMap[hours[i]]) - 1
			max = Max(max, count)
		} else {
			idxMap[hours[i]] = i
		}
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
