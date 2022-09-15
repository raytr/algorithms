package minimum_window_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, "BANC", minWindow("ADOBECODEBANC", "ABC"))
	assert.Equal(t, "a", minWindow("a", "a"))
}
