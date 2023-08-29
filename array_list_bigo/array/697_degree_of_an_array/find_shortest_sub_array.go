package degree_of_an_array

/*
	https://leetcode.com/problems/degree-of-an-array/

	1 <= length <= 50000
	0 <= val <= 49999

	map -> interate all elements -> with each loop, update the max, maxVal

	integrate all elements again to find the fist and end of maxVal's occurrences

	time complexity O(n)
	space complexity: O(n)
*/

func findShortestSubArray1(nums []int) int {
	maxFreq := 0
	maxVals := make([]int, 0)
	freqMap := make(map[int]int)

	for _, num := range nums { //O(n)
		freqMap[num]++
		if freqMap[num] > maxFreq {
			maxFreq = freqMap[num]
			maxVals = []int{num}
		} else if freqMap[num] == maxFreq {
			maxVals = append(maxVals, num)
		}
	}

	min := 490000
	for _, val := range maxVals { //O(
		dis := 0
		first, last := -1, -1
		for i, num := range nums {
			if num == val && first == -1 {
				first = i
			}

			if num == val && first != -1 {
				last = i
			}
		}

		if first != -1 && last == -1 {
			dis = 1
		} else {
			dis = last - first + 1
		}

		min = getMin(min, dis)
	}

	return min
}

func findShortestSubArray2(nums []int) int {
	left := make(map[int]int)
	right := make(map[int]int)
	freqMap := make(map[int]int)

	for i, num := range nums {
		if _, ok := left[num]; !ok {
			left[num] = i
		}
		right[num] = i
		freqMap[num]++
	}

	ans := len(nums)
	degree := 0
	for _, v := range freqMap {
		if v > degree {
			degree = v
		}
	}

	for num, freq := range freqMap {
		if freq == degree {
			ans = getMin(ans, right[num]-left[num]+1)
		}
	}

	return ans
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
