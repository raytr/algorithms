package missing_numbe

import "fmt"

/*
   clarify questions:
   what is the maximum and minimum of length of this array
   what is the maximum and minimum of number of this array

   brute force space complexity : O(n):
       create a existMap map[int]bool
       and push all nums in to map with existMap[num] = true

       append -1 to nums
       run i from 0 to n
       => if existMap[nums[i]] not exist => return i

*/

func missingNumberBruteForce(nums []int) int {
	n := len(nums)
	existMap := make(map[int]bool)

	//add to map
	for _, num := range nums {
		existMap[num] = true
	}

	nums = append(nums, -1)
	for i := 0; i <= n; i++ {
		if !existMap[i] {
			return i
		}
	}
	return -1
}

/*
Time complexity : O(n)
Space complexity: O(1)
*/
func missingNumberBitManipulation(nums []int) int {
	n := len(nums)
	missingNumber := n
	for i := 0; i < n; i++ {
		a := nums[i]
		b := i ^ a
		fmt.Println(b)
		missingNumber ^= i ^ nums[i]
	}
	return missingNumber
}
