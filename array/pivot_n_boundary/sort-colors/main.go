package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
}

func sortColors(nums []int) {
	boundary := 0
	pivot := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < pivot {
			nums[boundary], nums[i] = nums[i], nums[boundary]
			boundary++
		}
	}

	boundary = len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > pivot {
			nums[boundary], nums[i] = nums[i], nums[boundary]
			boundary--
		}
	}

	fmt.Println(nums)
}
