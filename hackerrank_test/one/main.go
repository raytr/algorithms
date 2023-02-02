package one

import (
	"strconv"
	"strings"
)

func ModifyString(str string) string {
	strings.TrimSpace(str)
	newStr := ""
	for _, c := range str {
		if _, err := strconv.Atoi(string(c)); err != nil {
			newStr += string(c)
		}
	}

	output := ""
	for i := len(newStr) - 1; i >= 0; i-- {
		output += string(newStr[i])
	}
	return output
}
