package decode_string

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeString(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expectation string
	}{
		{"3[a]2[bc]", "3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"3", "3", ""},
		{"3333", "3333", ""},
		{"leetcode", "leetcode", "leetcode"},
		{"100[leetcode]", "100[leetcode]", "leetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcode"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := decodeString(tt.input)
			require.Equal(t, tt.expectation, got)
		})
	}
}
