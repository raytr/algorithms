package find_all_numbers_disappeared_in_an_array

import "fmt"

/*
	problem: https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

   assum all in add 1

   1 2 3 4 //index + 1
   2 3 2 3

   swap 2 with element that has index is 2, until current element and the value at position we check equal
   3 2 2 3
   2 2 3 3 => the value of first element and the value of position 2 has same value
   0 2 3 3 => set the value of first element is 0

   check the next element is 2 => the value and index + 1 same value => do nothing and check next element

   ...
   the last elemtnt has value is 3 != 4 => check value of position 3 => is same => set the value at position 4 = 0
   0 2 3 0


   travesal a new array


*/

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		isMatch := true
		a := nums[i]
		b := nums[nums[i]-1]
		fmt.Println(a, b)
		c := nums[i] != i+1
		d := nums[i] != nums[nums[i]-1]
		fmt.Println(c, d)
		for nums[i] != i+1 && nums[i] != nums[nums[i]-1] { //[0] = 2; [3-1] = 2
			isMatch = false
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i] // 2 2 3 3
			if nums[i] == i+1 {                                 // 3 and 0 + 1
				isMatch = true
				break
			}
		}
		if isMatch == false {
			nums[i] = 0
		}
	}

	res := make([]int, 0, n)
	for i, num := range nums {
		if num == 0 {
			res = append(res, i+1)
		}
	}

	return res
}
