package minimum_number_of_flips_to_make_the_binary_string_alternating_ipt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"01001001101", 2},
	}

	for _, tt := range testCases {
		t.Run("case 1", func(t *testing.T) {
			require.Equal(t, tt.want, minFlips(tt.input))
		})
	}
}
