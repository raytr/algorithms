package anagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubAnagrams(t *testing.T) {
	testCases := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{"listen, silent, hello", []string{"listen", "silent", "hello"}, [][]string{{"listen", "silent"}, {"hello"}}},
	}

	for _, tt := range testCases {
		got := anagrams(tt.input)
		assert.Equal(t, tt.want, got)
	}
}
