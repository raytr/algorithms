package detect_capital

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	assert.Equal(t, true, detectCapitalUse("USA"))
	assert.Equal(t, false, detectCapitalUse("FlaG"))
	assert.Equal(t, true, detectCapitalUse("a"))
	assert.Equal(t, true, detectCapitalUse("Leetcode"))
}
