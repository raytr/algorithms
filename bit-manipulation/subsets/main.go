package main

import (
	"fmt"
	"math"
	"strconv"
)

//Problem: https://leetcode.com/problems/subsets/?fbclid=IwAR36lGWQEjXiAGh8ZNw0qrl5BrXlQdAxjyBchgdG1s9O7EA86Rp6vktJCxw

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

//Complexity: O(n^2)
func subsets(nums []int) [][]int {
	//create bit mask of index
	bitMasks := []uint64{}
	for i := 0; i < int(math.Pow(2, float64(len(nums)))); i++ {
		bit := decimal2Binary(i)
		bitMasks = append(bitMasks, bit)
	}

	output := [][]int{}
	for _, bitMark := range bitMasks {
		element := []int{}
		for i := 0; i < len(nums); i++ {
			b := getBit(bitMark, i)
			if b == 1 {
				element = append(element, nums[i])
			}
		}
		output = append(output, element)
	}

	return output
}

func decimal2Binary(decimal int) uint64 {
	bitStr := ""
	for decimal > 0 {
		remainder := decimal % 2
		bitStr = strconv.Itoa(remainder) + bitStr
		decimal = decimal / 2
	}
	binary, _ := strconv.ParseUint(bitStr, 2, 64)
	return binary
}

func getBit(x uint64, pos int) uint64 {
	return x >> pos & 1
}
