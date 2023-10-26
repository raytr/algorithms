package isomorphic_strings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsIsomorphic(t *testing.T) {
	testCases := []struct {
		name        string
		s           string
		t           string
		expectation bool
	}{
		{"s = egg, t = add", "egg", "add", true},
		{"s = foo, t = bar", "foo", "bar", false},
		{"s = paper, t = title", "paper", "title", true},
		{"s = badc, t = baba", "badc", "baba", false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectation, isIsomorphic(tt.s, tt.t))
		})
	}
}
