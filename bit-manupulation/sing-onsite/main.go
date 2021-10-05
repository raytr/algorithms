package main

import (
	"fmt"
	"math"
	"strconv"
)

//Problem: given an array A of distinct numbers, generate all subsets of A.
//Example: A = [1,2,3]
//Subsets = [[], [1] , [2] , [1,2] , [3] , [1,3] , [2,3] , [1,2,3]]

func main() {
	fmt.Println(getSubSet([]int{1, 2, 3}))
}

func getSubSet(arr []int) [][]int {

	//create bit mask
	bitMarks := getBitMaskOfLength(arr)

	output := [][]int{}
	for _, bitMark := range bitMarks {
		element := []int{}
		for i := 0; i < len(arr); i++ {
			b := getBit(bitMark, i)
			if b == 1 {
				element = append(element, arr[i])
			}
		}
		output = append(output, element)
	}

	return output
}

func getBitMaskOfLength(arr []int) []uint64 {
	bitMasks := []uint64{}
	for i := 0; i < int(math.Pow(2, float64(len(arr)))); i++ {
		bit := decimal2Binary(i)
		bitMasks = append(bitMasks, bit)
	}
	return bitMasks
}

func decimal2Binary(decimal int) uint64 {
	bitStr := ""
	for decimal > 0 {
		remainder := decimal % 2
		bitStr = strconv.Itoa(remainder) + bitStr
		decimal = decimal / 2
	}
	binary, _ := strconv.ParseUint(bitStr, 0, 64)
	return binary
}

func getBit(x uint64, pos int) uint64 {
	a := x >> pos
	b := a & 1
	return b
}
