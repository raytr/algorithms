package number_of_longest_increasing_subsequence

/*
		PROBLEM: https://leetcode.com/problems/number-of-longest-increasing-subsequence/


	   stage:
	       find f[i] ~ the lenght of longest subsequence that ended at i
	       count[i] ~ number of LIS that ended at i
	   relatition between sub and big:
	       f[i] = f[j] + 1
	   base case:
	       f[1] = 1


	       j    i
	      1. 3. 5. 4. 7

	      i == 1; j == 0
	       fx[0] == 1
	      fx[1] == 2
	      count[j] == 1
	      count[i] == 1

	      i == 2; j == 0
	      fx[0] == 1
	      fx[2] == 3
	      count[0] == 1
	      count[2] == 1

	      i == 2; j == 1
	      fx[1] == 2
	      fx[2] == 3
	      count[1] == 1
	      count[2] == count[2] + count[1] == 2 (fx[1] + 1 == fx[2])

	      i == 3; j== 0
	      fx[0] == 1
	      fx[3] == 1
	      count[0] == 1
	      count[3] == count[2] + count[1] == 2 (fx[1] + 1 == fx[2])
*/

func findNumberOfLIS(nums []int) int {
	n, maxLen := len(nums), 1
	fx := make(map[int]int)
	fx[0] = 1
	count := make(map[int]int) //number of LIS that ended at A[i]
	count[0] = 1

	for i := 1; i < n; i++ {
		fx[i] = 1
		count[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if fx[j]+1 > fx[i] {
					fx[i] = fx[j] + 1
					count[i] = count[j]
				} else if fx[i] == fx[j]+1 {
					count[i] += count[j]
				}
			}
		}
		if fx[i] > maxLen {
			maxLen = fx[i]
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		if fx[i] == maxLen {
			res += count[i]
		}
	}
	return res
}
