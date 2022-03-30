package missing_numbe

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolutionBruteForce(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		exp   int
	}{
		{"3, 0, 1", []int{3, 0, 1}, 2},
		{"0, 1", []int{0, 1}, 2},
		{"9, 6, 4, 2, 3, 5, 7, 0, 1", []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.exp, missingNumberBruteForce(tt.input))
		})
	}
}

func TestSolutionBitManipulation(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		exp   int
	}{
		{"3, 0, 1", []int{3, 0, 1}, 2},
		{"0, 1", []int{0, 1}, 2},
		{"9, 6, 4, 2, 3, 5, 7, 0, 1", []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.exp, missingNumberBitManipulation(tt.input))
		})
	}
}
