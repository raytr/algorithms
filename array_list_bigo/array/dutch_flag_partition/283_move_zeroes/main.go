package main

import "fmt"

func main() {
	args := []int{0, 1, 0, 3, 12}
	moveZeroes(args)
	fmt.Println(args)
}

func moveZeroes(nums []int) {
	boundary := 0
	pivot := 0

	for i, num := range nums {
		if num > pivot {
			nums[i], nums[boundary] = nums[boundary], nums[i]
			boundary++
		}
	}
}
