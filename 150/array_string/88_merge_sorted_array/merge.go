package merge_sorted_array

//problem: https://leetcode.com/problems/merge-sorted-array
// time complexity: O(m+n)
// space complexity: O(m+n)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			i++
		} else {
			frontHalf := nums1[:i]
			backHalf := append([]int{nums2[j]}, nums1[i:len(nums1)-1]...)
			nums1 = append(frontHalf, backHalf...)
			j++
			i++
			m++
		}
	}

	//check are all elements in nums2 have been added to nums1, if not, add
	for j < n {
		secondHalf := append([]int{nums2[j]}, nums1[i:len(nums1)-1]...)
		nums1 = append(nums1[:i], secondHalf...)
		j++
		i++
	}
}
