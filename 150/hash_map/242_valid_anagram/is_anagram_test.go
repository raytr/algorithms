package valid_anagram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	testcases := []struct {
		description string
		s           string
		t           string
		want        bool
	}{
		{"s = \"anagram\", t = \"nagaram\"", "anagram", "nagaram", true},
		{"s = \"rat\", t = \"car\"", "rat", "car", false},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, isAnagram(tc.s, tc.t))
		})
	}
}
