package maximum_length_of_repeated_subarray

/*

	https://leetcode.com/problems/maximum-length-of-repeated-subarray/

*/

func findLengthBruteForce(nums1 []int, nums2 []int) int {
	/*
		brute force
					   with each number in num1, i'm going to prepare with all number in num2
					   if i find a number that is the same with the chosen number in num1
					   -> i will increase curMax move pointer p1 to p2

							     p+4
					nums1 = [1,2,3,2,1]
				  			 i
					nums2 = [3,2,1,4,7]

					max = 3

			time complexity: O(n^2)
			space compexity: contain



	*/
	max := 0
	for pilot, _ := range nums1 {
		for i := 0; i < len(nums2); i++ {
			step := 0
			curMax := 0
			for nums1[pilot+step] == nums2[i+step] {
				curMax++
				step++
				if pilot+step == len(nums1) || i+step == len(nums2) {
					break
				}
			}
			max = getMax(max, curMax)
		}
	}

	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
