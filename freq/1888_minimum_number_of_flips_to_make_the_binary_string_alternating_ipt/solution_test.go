package minimum_number_of_flips_to_make_the_binary_string_alternating_ipt

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		input       string
		expectation int
	}{
		{"01001001101", 2},
	}

	for _, tt := range tests {
		t.Run("case 1", func(t *testing.T) {
			require.Equal(t, tt.expectation, minFlips(tt.input))
		})
	}
}
