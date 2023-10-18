package is_subsequence

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsSubsequence(t *testing.T) {
	testCases := []struct {
		name        string
		s           string
		t           string
		expectation bool
	}{
		{"s = abc, t = ahbgdc", "abc", "ahbgdc", true},
		{"s = axc, t = ahbgdc", "axc", "ahbgdc", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expectation, isSubsequence(tc.s, tc.t))
		})
	}
}
