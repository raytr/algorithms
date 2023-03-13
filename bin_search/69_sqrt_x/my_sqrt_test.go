package sqrt_x

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"4", 4, 2},
		{"8", 8, 2},
		{"0", 0, 0},
		{"1", 1, 1},
		{"2", 2, 1},
		{"2", 2, 1},
		{"3", 3, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mySqrt(tt.input))
		})
	}
}
