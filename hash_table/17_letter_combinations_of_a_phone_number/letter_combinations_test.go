package letter_combinations_of_a_phone_number

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLetterCombinations(t *testing.T) {
	testCases := []struct {
		name   string
		digits string
		want   []string
	}{
		{"23", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", "", []string{}},
		{"2", "2", []string{"a", "b", "c"}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.ElementsMatch(t, tt.want, letterCombinations(tt.digits))
		})
	}
}
