package minimum_insertion_steps_to_make_a_string_palindrome

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		exp  int
	}{
		{"zzazz", "zzazz", 0},
		{"mbadm", "mbadm", 2},
		{"leetcode", "leetcode", 5},
		{"tldjbqjdogipebqsohdypcxjqkrqltpgviqtqz", "tldjbqjdogipebqsohdypcxjqkrqltpgviqtqz", 25},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.exp, minInsertions(tt.s))
		})
	}
}
