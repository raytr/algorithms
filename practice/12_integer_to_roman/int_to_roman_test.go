package integer_to_roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRoman(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		want string
	}{
		{"num = 3", 3, "III"},
		{"num = 58", 58, "LVIII"},
		{"num = 1994", 1994, "MCMXCIV"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, intToRoman(tt.num))
		})
	}

}
