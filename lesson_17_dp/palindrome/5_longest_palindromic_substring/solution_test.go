package longest_palindromic_substring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	//assert.Equal(t, "bab", longestPalindrome("babad"))
	//assert.Equal(t, "bbb", longestPalindrome("bbbab"))
	//assert.Equal(t, "bb", longestPalindrome("bb"))
	assert.Equal(t, "bbb", longestPalindrome("bbb"))
	//assert.Equal(t, "bab", longestPalindrome("babad"))
	//assert.Equal(t, "bb", longestPalindrome("cbbd"))
}
