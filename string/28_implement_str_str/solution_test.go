package implement_str_str

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStrStr(t *testing.T) {
	testCases := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"mississippi", "issip", 4},
		{"aaaa", "aaaaa", -1},
		{"a", "a", 0},
	}

	for _, tt := range testCases {
		t.Run(tt.haystack, func(t *testing.T) {
			require.Equal(t, tt.want, strStr(tt.haystack, tt.needle))
		})
	}
}
