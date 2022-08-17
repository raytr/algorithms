package implement_str_str

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack    string
		needle      string
		expectation int
	}{
		{"mississippi", "issip", 4},
		{"aaaa", "aaaaa", -1},
		{"a", "a", 0},
	}

	for _, tt := range tests {
		t.Run(tt.haystack, func(t *testing.T) {
			require.Equal(t, tt.expectation, strStr(tt.haystack, tt.needle))
		})
	}
}
