package main

//describe: https://leetcode.com/problems/range-sum-query-2d-immutable/

import "fmt"

type NumMatrix struct {
	matrix [][]int
}

func main() {
	//input := [][]int{
	//	{3,0,1,4,2},
	//	{5,6,3,2,1},
	//	{1,2,0,1,5},
	//	{4,1,0,1,7},
	//	{1,3,0,0,5},
	//}
	//numMatrix := Constructor(input)
	//fmt.Println(numMatrix.SumRegion(2, 1, 4, 3))
	//
	//input2 := [][]int{{1}}
	//numMatrix2 := Constructor(input2)
	//fmt.Println(numMatrix2.SumRegion(0,0,0,0))

	input3 := [][]int{{-4, -5}}
	numMatrix3 := Constructor(input3)
	fmt.Println(numMatrix3.SumRegion(0, 0, 0, 0))
	fmt.Println(numMatrix3.SumRegion(0, 0, 0, 1))
	fmt.Println(numMatrix3.SumRegion(0, 1, 0, 1))
}

//Complexity: O(n^2)
func Constructor(matrix [][]int) NumMatrix {
	numMatrix := NumMatrix{matrix: [][]int{}}

	//edge cache r = 0
	numMatrix.matrix = append(numMatrix.matrix, []int{matrix[0][0]})
	for c := 1; c < len(matrix[0]); c++ {
		numMatrix.matrix[0] = append(numMatrix.matrix[0], numMatrix.matrix[0][c-1]+matrix[0][c])
	}

	for r := 1; r < len(matrix); r++ {
		numMatrix.matrix = append(numMatrix.matrix, []int{})
		numMatrix.matrix[r] = append(numMatrix.matrix[r], numMatrix.matrix[r-1][0]+matrix[r][0])
		for c := 1; c < len(matrix[0]); c++ {
			numMatrix.matrix[r] = append(numMatrix.matrix[r], numMatrix.matrix[r][c-1]+
				matrix[r][c]+
				numMatrix.matrix[r-1][c]-
				numMatrix.matrix[r-1][c-1])
		}
	}

	return numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 == 0 && row2 == 0 && col2 == 0 {
		return this.matrix[0][0]
	}

	sum := this.matrix[row2][col2]

	top := 0
	if row1 == 0 {
		top = this.matrix[0][col2]
	} else {
		top = this.matrix[row1-1][col2]
	}

	left := 0
	if col1 == 0 {
		left = this.matrix[row2][0]
	} else {
		left = this.matrix[row2][col1-1]
	}

	leftTop := 0
	if row1 == 0 && col1 == 0 {
		leftTop = this.matrix[0][0]
	} else if row1 == 0 {
		leftTop = this.matrix[0][col1-1]
	} else if col1 == 0 {
		leftTop = this.matrix[row1-1][0]
	} else {
		leftTop = this.matrix[row1-1][col1-1]
	}

	return sum - top - left + leftTop
}
