package palindrome_number

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input int
		exp   bool
	}{
		{"case: 121", 121, true},
		{"case: 1221", 1221, true},
		{"case: 0", 0, true},
		{"case: -121", -121, false},
		{"case: 10", 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, isPalindrome(tt.input), tt.exp)
		})
	}
}
