package roman_to_integer

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name           string
		expression     string
		expectedResult int
	}{
		{"happy case 1", "III", 3},
		{"happy case 2", "LVIII", 58},
		{"happy case 3", "MCMXCIV", 1994},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectedResult, romanToInt(tt.expression), tt.name)
		})
	}
}
