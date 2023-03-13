package valid_anagram

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsParagram(t *testing.T) {
	testCases := []struct {
		name  string
		input []string
		want  bool
	}{
		{"should be true", []string{"silent", "listen"}, true},
		{"should be false", []string{"hello", "hallo"}, false},
	}

	for _, tt := range testCases {
		require.Equal(t, tt.want, isAnagram(tt.input[0], tt.input[1]))
	}
}
