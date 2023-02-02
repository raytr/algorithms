package max_even_sum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMaximumEvenSum(t *testing.T) {
	tests := []struct {
		name string
		val  []int32
		want int64
	}{
		{"[6, 3, 4, -1, 9, 17]", []int32{6, 3, 4, -1, 9, 17}, 38},
		{"[2, 3, 6, -5, 10, 1, 1]", []int32{2, 3, 6, -5, 10, 1, 1}, 22},
		{"[-1, -2, -3, 8, 7]", []int32{-1, -2, -3, 8, 7}, 14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, getMaximumEvenSum(tt.val))
		})
	}
}
