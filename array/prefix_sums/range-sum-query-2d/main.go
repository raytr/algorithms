package main

import "fmt"

type NumMatrix struct {
	matrix [][]int
}

func main() {
	fmt.Println(1)
	//input := [][]int{
	//	{3,0,1,4,2},
	//	{5,6,3,2,1},
	//	{1,2,0,1,5},
	//	{4,1,0,1,7},
	//	{1,3,0,3,5},
	//}
	//numMatrix := Constructor(input)
	//fmt.Println(numMatrix.SumRegion(1,1,4,4))
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{
		matrix: matrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for r := row1 - 1; r < row2-1; r++ {
		for c := col1 - 1; c < col2-1; c++ {
			sum += this.matrix[r][c]
		}
	}
	return sum
}
