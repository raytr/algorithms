package subarray_sum_equal_k

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func TestSubArraySumEqualK(t *testing.T) {
	tests := []struct {
		want int
		arr  []int
		k    int
		name string
	}{
		{2, []int{5, -1, 3, 2, 9, 4}, 5, "wrong at first"},
		{3, []int{2, 5, 0, 2, 5}, 5, "wrong at second"},
		{2, []int{1, 1, 1}, 2, "wrong at third"},
		{4, []int{1, 2, 1, 2, 1}, 3, "wrong at forth"},
		{1, []int{-1, -1, 1}, 0, "wrong at fifth"},
		{3, []int{1, -1, 0}, 0, "wrong at sixth"},
		{2, []int{1, 2, 3}, 3, "wrong at seventh"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, subarraySum(tt.arr, tt.k))
		})
	}
}
