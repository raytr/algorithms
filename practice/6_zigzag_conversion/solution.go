package zigzag_conversion

/*

	problem: https://leetcode.com/problems/zigzag-conversion/

	   create 2D array has numRows subarray
	   r, i = 0
	   while with condition is i < n (n is len of s)
	       while r < numRows && i < n
	           Rows[r] = append(s[i])
	           r++; i++
	       while r >= 0 && i < n
	           Rows[r] = append(s[i])
	           r--; i++
	       r+= 2

	   get each character in rows and push it to res
	   return res

	   time: O(n)
	   space: O(n) //because created array 2D to store characters
*/

func convert(s string, numRows int) string {
	//edge cases
	if numRows == 1 {
		return s
	}

	rows := make([][]string, 0, numRows)
	for i := 0; i < numRows; i++ {
		rows = append(rows, make([]string, 0, len(s)))
	}

	i, r, n := 0, 0, len(s)
	for i < n {
		for r < numRows && i < n {
			rows[r] = append(rows[r], string(s[i]))
			r++
			i++
		}
		r -= 2 // ex: numRows = 3 => move to center (4 -> 2)
		for r >= 0 && i < n {
			rows[r] = append(rows[r], string(s[i]))
			r--
			i++
		}
		r += 2
	}

	res := ""
	for _, row := range rows {
		for _, c := range row {
			res += c
		}
	}
	return res
}
