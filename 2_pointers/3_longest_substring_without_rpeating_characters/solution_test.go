package longest_substring_without_rpeating_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"), "abcabcbb")
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"), "bbbbb")
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"), "pwwkew")
	assert.Equal(t, 0, lengthOfLongestSubstring(""))
	assert.Equal(t, 1, lengthOfLongestSubstring(" "))
}
