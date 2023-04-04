package reverse_integer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		name string
		x    int
		want int
	}{
		{"123", 123, 321},
		{"-123", -123, -321},
		{"1534236469", 1534236469, 0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, reverse(tt.x))
		})
	}
}
