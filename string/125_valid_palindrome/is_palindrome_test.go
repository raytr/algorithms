package valid_palindrome

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want bool
	}{
		{"A man, a plan, a canal: Panama", "A man, a plan, a canal: Panama", true},
		{"race a car", "race a car", false},
		{"0P", "0P", false},
		{"\",M 9y\\\"yj\\\"j9 M,\"", "\",M 9y\\\"yj\\\"j9 M,\"", false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, isPalindrome(tt.s))
		})
	}
}
