package excel_sheet_column_number

/*
	problem: https://leetcode.com/problems/excel-sheet-column-number/
   do columnTitle just contains lower leters, do it contain number, uppwer letter, ...?. Rep: consists only of uppercase English letters.
   what is max & min of columnTitle length?  Rep: 1 <= columnTitle.length <= 7
   what is range of columnTitle?  Rep: ["A", "FXSHRXW"].

   solution:
       make new dictionary is map of [character]int
       read from right -> left. Move by one, *28

   time complexity: O(n). n is length of columnTitle
*/

func titleToNumber(columnTitle string) int {
	dic := newDictionary()
	n := len(columnTitle)
	sum := 0
	coefficient := 1

	for i := n - 1; i >= 0; i-- {
		char := columnTitle[i]
		num := dic[char] * coefficient
		sum += num
		coefficient *= 26
	}

	return sum
}

func newDictionary() map[uint8]int {
	dic := make(map[uint8]int)
	dic['A'] = 1
	dic['B'] = 2
	dic['C'] = 3
	dic['D'] = 4
	dic['E'] = 5
	dic['F'] = 6
	dic['G'] = 7
	dic['H'] = 8
	dic['I'] = 9
	dic['J'] = 10
	dic['K'] = 11
	dic['L'] = 12
	dic['M'] = 13
	dic['N'] = 14
	dic['O'] = 15
	dic['P'] = 16
	dic['Q'] = 17
	dic['R'] = 18
	dic['S'] = 19
	dic['T'] = 20
	dic['U'] = 21
	dic['V'] = 22
	dic['W'] = 23
	dic['X'] = 24
	dic['Y'] = 25
	dic['Z'] = 26

	return dic
}
