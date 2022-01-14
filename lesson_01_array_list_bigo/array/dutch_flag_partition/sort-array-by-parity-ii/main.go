package main

//problem: https://leetcode.com/problems/sort-array-by-parity-ii/

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{888, 505, 627, 846}))
	fmt.Println(sortArrayByParityII([]int{3, 1, 4, 2}))
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
	fmt.Println(sortArrayByParityII([]int{3, 0, 4, 0, 2, 1, 3, 1, 3, 4}))
	fmt.Println(sortArrayByParityII([]int{648, 831, 560, 986, 192, 424, 997, 829, 897, 843}))
	fmt.Println(sortArrayByParityII([]int{0, 3, 4, 0, 2, 1, 3, 1, 3, 4}))
	fmt.Println(sortArrayByParityII([]int{0, 1, 2, 4, 6, 8, 3, 5, 7, 9}))
}

func sortArrayByParityII(nums []int) []int {
	boundary := 0
	remainder := 0
	i := 0
	for i < len(nums) {
		if !((nums[i]%2 == 0 && i%2 == 0) || (nums[i]%2 == 1 && i%2 == 1)) {
			if boundary == i {
				remainder = i % 2
				i++
			} else if i%2 != remainder {
				nums[boundary], nums[i] = nums[i], nums[boundary]
				boundary = +1
				i = boundary
			} else {
				i++
			}
		} else if i == boundary {
			i++
			boundary++
		} else {
			i++
		}
	}

	return nums
}
