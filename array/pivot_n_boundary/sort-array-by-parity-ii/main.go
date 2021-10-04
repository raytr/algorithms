package main

import "fmt"

func main() {
	//fmt.Println(sortArrayByParityII([]int{888,505,627,846}))
	//fmt.Println(sortArrayByParityII([]int{3,1,4,2}))
	//fmt.Println(sortArrayByParityII([]int{3,0,4,0,2,1,3,1,3,4}))
	//fmt.Println(sortArrayByParityII([]int{648,831,560,986,192,424,997,829,897,843}))
	fmt.Println(sortArrayByParityII([]int{0, 3, 4, 0, 2, 1, 3, 1, 3, 4}))
}

func sortArrayByParityII(nums []int) []int {

	boundary := 0
	for i := 0; i < len(nums); i++ {
		if !(i%2 == 0 && nums[i]%2 == 0) && !(i%2 == 1 && nums[i]%2 == 1) {
			nums[i], nums[boundary] = nums[boundary], nums[i]
			boundary = i + 1
		}
	}

	boundary = len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if !(i%2 == 0 && nums[i]%2 == 0) && !(i%2 == 1 && nums[i]%2 == 1) {
			nums[i], nums[boundary] = nums[boundary], nums[i]
			boundary = i - 1
		}
	}

	return nums
}
