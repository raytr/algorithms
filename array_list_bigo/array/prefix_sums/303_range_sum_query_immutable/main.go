package _03_range_sum_query_immutable

//describe: https://leetcode.com/problems/range-sum-query-immutable/

type NumArray struct {
	nums []int
}

// Complexity: O(n)
func Constructor(nums []int) NumArray {
	numArray := NumArray{nums: []int{}}
	numArray.nums = append(numArray.nums, nums[0])
	for i := 1; i < len(nums); i++ {
		prefixSum := nums[i] + numArray.nums[i-1]
		numArray.nums = append(numArray.nums, prefixSum)
	}
	return numArray
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.nums[right]
	}
	return this.nums[right] - this.nums[left-1]
}
