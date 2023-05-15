package spiral_matrix

/*
	https://leetcode.com/problems/spiral-matrix/

	design direction:
		moveRight: [0, 1]
		moveDownward: [1, 0]
		moveLeft: [-1, 0]
		moveUpward: [-1, 0]

	move right until meet maxWid
	move downward until meed maxHeight
	move left until meet minWid
	move upward until meet minHeight

	because through each element once -> time complexity í: O(n)
*/

func spiralOrder(matrix [][]int) []int {
	moveRight := []int{0, 1}
	moveLeft := []int{-0, -1}
	moveDownward := []int{1, 0}
	moveUpward := []int{-1, 0}

	minWid := 0
	maxWid := len(matrix[0]) - 1
	minHeight := 0
	maxHeight := len(matrix) - 1
	total := len(matrix) * len(matrix[0])
	result := make([]int, 0, total)
	cur := []int{0, -1} //start from outside
	count := 0

	for count < total {

		for i := minWid; i <= maxWid; i++ {
			cur[0] += moveRight[0]
			cur[1] += moveRight[1]
			result = append(result, matrix[cur[0]][cur[1]])
			count++
			if count == total {
				return result
			}
		}

		for i := minHeight + 1; i <= maxHeight; i++ {
			cur[0] += moveDownward[0]
			cur[1] += moveDownward[1]
			result = append(result, matrix[cur[0]][cur[1]])
			count++
			if count == total {
				return result
			}
		}

		for i := maxWid; i >= minWid+1; i-- {
			cur[0] += moveLeft[0]
			cur[1] += moveLeft[1]
			result = append(result, matrix[cur[0]][cur[1]])
			count++
			if count == total {
				return result
			}
		}

		for i := maxHeight; i >= minHeight+2; i-- {
			cur[0] += moveUpward[0]
			cur[1] += moveUpward[1]
			result = append(result, matrix[cur[0]][cur[1]])
			count++
			if count == total {
				return result
			}
		}

		minWid++
		maxWid--
		minHeight++
		maxHeight--
	}

	return result
}
