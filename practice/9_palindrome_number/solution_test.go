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
		{"121", 121, true},
		{"1221", 1221, true},
		{"0", 0, true},
		{"-121", -121, false},
		{"10", 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, isPalindrome(tt.input), tt.exp)
		})
	}
}
