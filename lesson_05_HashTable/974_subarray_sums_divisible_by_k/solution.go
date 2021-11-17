//package 974_subarray_sums_divisible_by_k
package main

import "fmt"

func main() {
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
	fmt.Println(subarraysDivByK([]int{-5}, 5))
	fmt.Println(subarraysDivByK([]int{-1, 2, 9}, 2))
	fmt.Println(subarraysDivByK([]int{2, -2, 2, -4}, 6))
}

// problem: https://leetcode.com/problems/subarray-sums-divisible-by-k/

/*
	 A % K = B % K
	=> A - B
	If a % k = x and b % k = x => (a - b) % k == 0

	If we have:  a b c =>  ab,  bc,  ac = 3 results.
	=> formula: (f - 1) * f / 2
	F is frequency, the time this number % k == x


	index:            0 ,     1 ,      2,        3,      4,      5        6
	Mod               2       1        2         3       2       1        2

	Freq[2] = 		  1                2                 3                 4
	Subarray  =       0                1                 2                 3

	Freq[1] =                  1                                  2
	Subarray =               0                                    1

	Freq[3] =                                        1
	Subarray =                                     0

*/
func subarraysDivByK(nums []int, k int) int {
	prefixSums := make([]int, 0, len(nums))
	prefixSums = append(prefixSums, nums[0])
	for i := 1; i < len(nums); i++ {
		prefixSums = append(prefixSums, prefixSums[i-1]+nums[i])
	}

	freq := make(map[int]int)
	for _, pfs := range prefixSums {
		remainder := getRemainder(pfs, k)
		freq[remainder]++
	}

	count := freq[0]
	for _, v := range freq {
		if v > 0 {
			count = count + (v-1)*v/2
		}
	}

	return count
}

func getRemainder(a, k int) int {
	return (a%k + k) % k
}
