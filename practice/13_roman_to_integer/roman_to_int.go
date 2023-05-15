package roman_to_integer

/*

I             1
V             5
X             10
L             50
C             100
D             500
M             1000

	XIV -> 14
	X -> sum = 10

	if len == 1 => return dict[s[0]]

	if i != len-1 && v[i+1] > v[i] => sum -= v[i]
	else sum +=

	return sum


    time complexity is O(1)
    space complexity is O(1)


*/

func romanToInt(s string) int {
	dict := buildDict()
	sum := 0

	for i := 0; i < len(s); i++ {
		if i <= len(s)-2 && dict[s[i]] < dict[s[i+1]] {
			sum -= dict[s[i]]
		} else {
			sum += dict[s[i]]
		}
	}
	return sum
}

func buildDict() map[byte]int {
	dict := make(map[byte]int)
	dict['I'] = 1
	dict['V'] = 5
	dict['X'] = 10
	dict['L'] = 50
	dict['C'] = 100
	dict['D'] = 500
	dict['M'] = 1000
	return dict
}
