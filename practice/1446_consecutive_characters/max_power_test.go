package consecutive_characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPower(t *testing.T) {
	tests := []struct {
		name        string
		s           string
		expectation int
	}{
		{"s = leetcode", "leetcode", 2},
		{"s = abbcccddddeeeeedcba", "abbcccddddeeeeedcba", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation, maxPower(tt.s))
		})
	}
}
