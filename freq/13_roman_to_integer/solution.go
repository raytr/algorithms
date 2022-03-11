package roman_to_integer

/*

I             1
V             5
X             10
L             50
C             100
D             500
M             1000

    XIII => 8
    X => 5 => sum += 5
    I => 1  => sum += 1

    sum == 8

    "LVIII"
    L => 50 => sum+=50
    V => 5 => sum+=5
    + 1, + 1. +1 => 58
    ....

    time complexity is O(1)
    space complexity is O(1)



*/

func romanToInt(s string) int {
	dict := buildDict()
	sum := 0

	for i := 0; i < len(s); i++ {
		if i != len(s)-1 && dict[s[i]] < dict[s[i+1]] {
			sum -= dict[(s[i])]
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
