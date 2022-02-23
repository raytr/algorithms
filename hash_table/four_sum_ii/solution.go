package four_sum_ii

//problem: https://leetcode.com/problems/4sum-ii/

/*
	separate become to part: nums1, nums2 and nums3, nums4
	use nesting for loop to find all combination cases

	with first loop (nums1, nums2)
	get sum and save sums to map. map has key is sum and value is the time that sum appears

	with second loop (sums3, nums4)
	with each iterate, calculate sum and check negative of sum is exist in map:
		- if exist => count + value of map (time that sum appear)

	Ex: [1,0] [0,1] => 3 combition:
		map[0] = 1
		map[1] = 2
		map[2] = 1

		 => need to find -1, 2 , 0 in second part
*/

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	combitionMap := make(map[int]int)
	count := 0
	for _, i := range nums1 {
		for _, j := range nums2 {
			sum := i + j
			if _, ok := combitionMap[sum]; !ok {
				combitionMap[sum] = 1
			} else {
				combitionMap[sum] += 1
			}
		}
	}

	for _, k := range nums3 {
		for _, l := range nums4 {
			sum := k + l
			if _, exist := combitionMap[-sum]; exist {
				count = count + combitionMap[-sum]
			}
		}
	}

	return count
}
