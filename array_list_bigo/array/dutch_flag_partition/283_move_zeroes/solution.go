package main

import "fmt"

func main() {
	args := []int{0, 1, 0, 3, 12}
	moveZeroes(args)
	fmt.Println(args)
}

func moveZeroes(nums []int) {
	boundary := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[boundary] {
			nums[i], nums[boundary] = nums[boundary], nums[i]
		}
	}
}
