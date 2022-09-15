package longest_substring_with_at_most_k_distinct_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, 0, lengthOfLongestSubstringKDistinct("aaa", 0))
}
