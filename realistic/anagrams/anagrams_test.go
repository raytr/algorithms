package anagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubAnagrams(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{"listen, silent, hello", []string{"listen", "silent", "hello"}, [][]string{{"listen", "silent"}, {"hello"}}},
	}

	for _, tt := range tests {
		got := anagrams(tt.input)
		assert.Equal(t, tt.want, got)
	}
}
