package longest_word_in_dictionary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, "world", longestWord([]string{"w", "wo", "wor", "worl", "world"}))
	assert.Equal(t, "apple", longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
}
