package excel_sheet_column_title

/*
	https://leetcode.com/problems/excel-sheet-column-title/

	28 / 26 == 1 => A
	28 % 26 == 2 => B
=> AB

	730 / 26 = 28, remainder 2(B)
	28 / 26 = 1, remainder 2(B)

	ABB

	for columnNumber > 26
		remainder = columnNumber % 26 => add string(remainder) to front of res
		columnNumber /= 26

	if coumnNum > 0 => add string(columnNumber) to front of res

	return front
*/

func convertToTitle(columnNumber int) string {
	var res string
	for columnNumber >= 26 {
		remainder := columnNumber % 26

		if remainder == 0 {
			res = "Z" + res
			columnNumber = columnNumber/26 - 1
		} else {
			res = string(rune(64+remainder)) + res
			columnNumber /= 26
		}
	}

	if columnNumber > 0 {
		res = string(rune(64+columnNumber)) + res
	}

	return res
}
