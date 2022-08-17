package maximum_subarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaximumSubarray(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"[-2, 1, -3, 4, -1, 2, 1, -5, 4]", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"[1]", []int{1}, 1},
		{"[5,4,-1,7,8]", []int{5, 4, -1, 7, 8}, 23},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, maxSubArray(tt.input))
		})
	}
}
