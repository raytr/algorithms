package longest_consecutive_sequence

/*
	problem: https://leetcode.com/problems/longest-consecutive-sequence/description/

	[100,4,200,1,3,2] => put all to map => O(n)
	numMap[100]=true
	numMap[4]=true
	numMap[200]=true
	numMap[1]=true
	numMap[3]=true
	numMap[2]=true

	100 -> find 101 in map  => O(1)
	if exsit -> upcate count++, continue
		update max

	time complexity: O(n^2)
*/

func longestConsecutive(nums []int) int {
	max := 0
	dictMap := make(map[int]bool)
	for _, num := range nums { //O(n)
		dictMap[num] = true
	}

	for num, _ := range dictMap { //O(n)

		//check is exist number front of current num -> to iterate each number once
		if dictMap[num-1] == false {
			count := 1
			i := 1
			for dictMap[num+i] == true {
				count++
			}
			max = getMax(max, count)
		}
	}

	return max
}

func longestConsecutiveBruteForce(nums []int) int {
	max := 0
	dict := make(map[int]bool)
	//O(n)
	for _, num := range nums {
		dict[num] = true
	}

	//O(n)
	for _, num := range nums {
		count := 1
		i := 1
		for dict[num+i] == true { //O(n)
			count++
			i++
		}

		max = getMax(max, count)
	}

	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
