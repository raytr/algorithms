package missing_number

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolutionBruteForce(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		exp   int
	}{
		{"3, 0, 1", []int{3, 0, 1}, 2},
		{"0, 1", []int{0, 1}, 2},
		{"9, 6, 4, 2, 3, 5, 7, 0, 1", []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.exp, missingNumberBruteForce(tt.input))
			require.Equal(t, tt.exp, missingNumberBitManipulation(tt.input))
		})
	}
}
