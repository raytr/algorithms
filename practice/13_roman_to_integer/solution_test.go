package roman_to_integer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name           string
		expression     string
		expectedResult int
	}{
		{"happy case 1", "III", 3},
		{"happy case 2", "LVIII", 58},
		{"happy case 3", "MCMXCIV", 1994},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectedResult, romanToInt(tt.expression), tt.name)
		})
	}
}
