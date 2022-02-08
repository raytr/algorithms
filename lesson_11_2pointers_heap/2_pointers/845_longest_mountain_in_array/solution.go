package longest_mountain_in_array

//problem: https://leetcode.com/problems/longest-mountain-in-array/

func longestMountain(arr []int) int {
	/*
		   we have 2 pointers up & down
		   up will move far as possible, when meet peak
		   down = peak and move down pointer far as possible

		   if up != peak != down => we have a mountain => down - up + 1
		   if up == peak => not up yet => move up continues

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

//
///*
//
//we add max(int) at last array
//we have 2 pointer f and s = 0,
//while f < n - 1
//    while arr[f] < [f+1] => f++
//    peak := f
//    while arr[f] > [f+1] => f++
//    if peak != f
//        longest = f - s
//        s = f
//
//    f++
//    s++
//*/
//
//func longestMountain(arr []int) int {
//	if len(arr) < 3 {
//		return 0
//	}
//
//	longest := 0
//	n, s, f := 0, 0, len(arr)
//	arr = append(arr, math.Int32.Max)
//
//	for f < n {
//		for arr[f] < arr[f+1] {
//			f++
//		}
//		peak := f
//		for arr[f] > arr[f+1] {
//			f++
//		}
//		if s < peak && f > peak {
//			longest = f - s
//			s = f
//		}
//		s++
//		f++
//	}
//}
