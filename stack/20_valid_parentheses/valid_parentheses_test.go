package valid_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidParentheses(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{"s = ()", "()", true},
		{"s = ()[]{}", "()[]{}", true},
		{"s = (]", "(]", false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isValid(tt.input))
		})
	}
}
