package product_of_array_except_self

/*
	problem: https://leetcode.com/problems/product-of-array-except-self

	result pointer is an array

	put 1 to result (because this is multiple)
	i have left pointer is 1, move left pointer from left -> right.
		result[left] = result[left-1] * nums[left-1]

	right pointer is len[nums]-1; move right pointer from right to left
		result[right-1] = result[right - 1] + nums[right]

	returen result

*/

func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n, n)
	right := make([]int, n, n)

	// L[i] contains the product of all the elements to the left
	// Note: for the element at index '0', there are no elements to the left,
	// so L[0] would be 1
	left[0] = 1
	for i := 1; i < n; i++ {
		// L[i - 1] already contains the product of elements to the left of 'i - 1'
		// Simply multiplying it with nums[i - 1] would give the product of all
		// elements to the left of index 'i'
		left[i] = left[i-1] * nums[i-1]
	}

	// R[i] contains the product of all the elements to the right
	// Note: for the element at index 'length - 1', there are no elements to the right,
	// so the R[length - 1] would be 1
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		// R[i + 1] already contains the product of elements to the right of 'i + 1'
		// Simply multiplying it with nums[i + 1] would give the product of all
		// elements to the right of index 'i'
		right[i] = right[i+1] * nums[i+1]
	}

	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		// For the first element, R[i] would be product except self
		// For the last element of the array, product except self would be L[i]
		// Else, multiple product of all elements to the left and to the right
		result = append(result, left[i]*right[i])
	}

	return result
}
