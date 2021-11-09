package longest_mountain_in_array

//problem: https://leetcode.com/problems/longest-mountain-in-array/

func longestMountain(arr []int) int {
	/*
		   we have 2 pointers up & down
		   up will move farest as possible, when meet peak
		   down = peak and move down pointer farest as possible

		   if up != peak != down => we have a mountain => down - up + 1
		   if up == peak => not up yet => move up continuos

			complexity: O(n)
	*/

	max, i := 0, 0
	n := len(arr)
	for i < n {
		up := i
		for up+1 < n && arr[up+1] > arr[up] {
			up++
		}
		down := up
		for down+1 < n && arr[down+1] < arr[down] {
			down++
		}
		if i != up && up != down {
			// we have a mountain has long:
			max = Max(max, down-i+1)
			i = down
		} else {
			i++
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
