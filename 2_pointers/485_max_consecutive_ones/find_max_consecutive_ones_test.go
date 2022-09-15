package max_consecutive_ones

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"[1,1,0,1,1,1]", []int{1, 1, 0, 1, 1, 1}, 3},
		{"[1,0,1,1,0,1]", []int{1, 0, 1, 1, 0, 1}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, findMaxConsecutiveOnes(tt.nums))
		})
	}
}
