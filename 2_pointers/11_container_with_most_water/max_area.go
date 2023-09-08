package container_with_most_water

/*
	https://leetcode.com/problems/container-with-most-water/
	 l
	[1,8,6,2,5,4,8,3,7]
                     r

	max = 0
	compare l & r, if value at l smaller than r => move l to right => update max. And vise versa
	util l >= r

	time complexity: O(n)

*/

func maxArea(height []int) int {
	n := len(height)
	max := 0
	l, r := 0, n-1

	for l < r {
		container := 0
		if height[l] == height[r] {
			container = (r - l) * height[l]
			l++
			r--
		} else if height[l] < height[r] {
			container = (r - l) * height[l]
			l++
		} else {
			container = (r - l) * height[r]
			r--
		}
		max = getMax(max, container)
	}

	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
