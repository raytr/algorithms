package reverse_integer

import (
	"fmt"
	"math"
	"strconv"
)

/*
	https://leetcode.com/problems/reverse-integer/

   convert this to string,
   check the first character, if it is "-" => negative == true and remove the first character

   for i to i < n / 2
       swap(str[i], str[n/2-1])

   O(logN)
*/

func reverse(x int) int {
	str := fmt.Sprintf("%v", x)
	negative := false
	if string(str[0]) == "-" {
		negative = true
		str = str[1:]
	}

	r := []rune(str)
	n := len(str)
	for i := 0; i < n/2; i++ {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}

	//rune to string
	res, _ := strconv.Atoi(string(r))
	if negative {
		res = res * -1
	}

	if float64(res) < math.Pow(float64(-2), float64(31)) || float64(res) > math.Pow(float64(2), float64(31))-1 {
		return 0
	}
	return res
}
