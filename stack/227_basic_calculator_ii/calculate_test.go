package basic_calculator_ii

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculate(t *testing.T) {
	testCases := []struct {
		description string
		s           string
		want        int
	}{
		{"s=3+2*2", "3+2*2", 7},
		{"s= 3/2", " 3/2", 1},
		{"s= 3+5 / 2 ", " 3+5 / 2 ", 5},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			require.Equal(t, testCase.want, calculate(testCase.s))
		})
	}
}
