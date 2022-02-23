package longest_substring_with_at_most_k_distinct_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 0, lengthOfLongestSubstringKDistinct("aaa", 0))
}
