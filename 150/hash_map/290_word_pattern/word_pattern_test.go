package word_pattern

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordPattern(t *testing.T) {
	testCases := []struct {
		description string
		pattern     string
		s           string
		expectation bool
	}{
		{"pattern = abba, s = dog cat cat dog", "abba", "dog cat cat dog", true},
		{"pattern = abba, s = dog cat cat fish", "abba", "dog cat cat fish", false},
		{"pattern = abba, s = dog dog dog dog", "abba", "dog dog dog dog", false},
		{"pattern = aaaa, s = dog cat cat dog", "aaaa", "dog cat cat dog", false},
		{"pattern = aaa, s = aa aa aa aa", "aaa", "aa aa aa aa", false},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			require.Equal(t, tt.expectation, wordPattern(tt.pattern, tt.s))
		})
	}
}
