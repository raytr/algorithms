package roman_to_integer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		name           string
		expression     string
		expectedResult int
	}{
		{"case V", "V", 5},
		{"case III", "III", 3},
		{"case LVIII", "LVIII", 58},
		{"case MCMXCIV", "MCMXCIV", 1994},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectedResult, romanToInt(tt.expression), tt.name)
		})
	}
}
