package search_a_2d_matrix_ii

/*
	problem: https://leetcode.com/problems/search-a-2d-matrix-ii/

   begin from bot left => c = 0; r = len(r)-1
   while c < len(matrix[0]) - 1 && r >= 0
       if matrix[r][c] == target => return true
       if matrix[r][c] < target => c--
       ortherwise r++

   return false

	complexity: O(n+m)
*/

func searchMatrix(matrix [][]int, target int) bool {

	c, r := 0, len(matrix)-1
	for c < len(matrix[0]) && r >= 0 {
		if matrix[r][c] == target {
			return true
		}
		if matrix[r][c] < target {
			c++
		} else {
			r--
		}
	}
	return false
}
