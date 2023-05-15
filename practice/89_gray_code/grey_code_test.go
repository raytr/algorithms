package gray_code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreyCode(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		expectation []int
	}{
		{"n = 2]", 2, []int{0, 1, 3, 2}},
		{"n = 1]", 1, []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation, grayCode(tt.n))
		})
	}
}
