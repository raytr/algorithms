package _83_ransom_note

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanConstruct(t *testing.T) {
	testCases := []struct {
		name        string
		ransomNote  string
		magazine    string
		expectation bool
	}{
		{"ransomNote = a, magazine = b", "a", "b", false},
		{"ransomNote = aa, magazine = ab", "aa", "ab", false},
		{"ransomNote = aa, magazine = aab", "aa", "aab", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectation, canConstruct(tt.ransomNote, tt.magazine))
		})
	}
}
