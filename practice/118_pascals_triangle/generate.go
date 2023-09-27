package pascals_triangle

/*

	problem: https://leetcode.com/problems/pascals-triangle/
	time complexity: O(numRows ^ 2)

*/

func generate(numRows int) [][]int {
	result := make([][]int, 0, numRows)
	if numRows >= 1 {
		result = append(result, []int{1})
	}

	if numRows >= 2 {
		result = append(result, []int{1, 1})
	}

	if numRows >= 3 {
		for i := 3; i <= numRows; i++ {
			//pop the last array
			lastArr := result[len(result)-1]
			level := []int{1}
			for j := 0; j < len(lastArr)-1; j++ {
				sum := lastArr[j] + lastArr[j+1]
				level = append(level, sum)
			}
			level = append(level, 1)
			result = append(result, level)
		}
	}

	return result
}
