package _52_peak_index_in_a_mountain_array

/*

problem: https://leetcode.com/problems/peak-index-in-a-mountain-array/

*/

func peakIndexInMountainArray(arr []int) int {
	lo := 0
	hi := len(arr) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}
		if mid == 0 || arr[mid] < arr[mid+1] {
			lo = mid + 1
		} else if mid == hi || arr[mid] < arr[mid-1] {
			hi = mid - 1
		}
	}
	return -1
}
