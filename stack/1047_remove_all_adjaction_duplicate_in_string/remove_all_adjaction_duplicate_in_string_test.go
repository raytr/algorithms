package remove_all_adjaction_duplicate_in_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	assert.Equal(t, "azxzy", removeDuplicates("azxxxzy"))

	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{"s= abbaca", "abbaca", "ca"},
		{"s= azxxzy", "azxxzy", "ay"},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeDuplicates(tt.input))
		})
	}
}
