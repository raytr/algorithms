package integer_to_roman

/*
	https://leetcode.com/problems/integer-to-roman/

	XL   X   IX   V   IV   I
	40   10  9    5   4    1

		 13           3
*/

type KeyValue struct {
	value int
	roman string
}

func intToRoman(num int) string {
	dict := buildDict()
	reusult := ""
	for i, kv := range dict {
		for num >= dict[i].value {
			reusult += kv.roman //X
			num -= kv.value     //4
		}
	}

	return reusult
}

func buildDict() []*KeyValue {
	// create dict follow descending order
	dict := make([]*KeyValue, 0, 13)
	dict = append(dict, &KeyValue{1000, "M"})
	dict = append(dict, &KeyValue{900, "CM"})
	dict = append(dict, &KeyValue{500, "D"})
	dict = append(dict, &KeyValue{400, "CD"})
	dict = append(dict, &KeyValue{100, "C"})
	dict = append(dict, &KeyValue{90, "XC"})
	dict = append(dict, &KeyValue{50, "L"})
	dict = append(dict, &KeyValue{40, "XL"})
	dict = append(dict, &KeyValue{10, "X"})
	dict = append(dict, &KeyValue{9, "IX"})
	dict = append(dict, &KeyValue{5, "V"})
	dict = append(dict, &KeyValue{4, "IV"})
	dict = append(dict, &KeyValue{1, "I"})

	return dict
}
