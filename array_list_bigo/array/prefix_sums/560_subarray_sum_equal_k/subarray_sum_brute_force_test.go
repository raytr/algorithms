package subarray_sum_equal_k

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func TestSubArraySumEqualK(t *testing.T) {
	testCases := []struct {
		name        string
		nums        []int
		k           int
		expectation int
	}{
		{"nums = [1,1,1], k = 2", []int{5, -1, 3, 2, 9, 4}, 5, 2},
		{"nums = [2, 5, 0, 2, 5], k = 5", []int{2, 5, 0, 2, 5}, 5, 3},
		{"nums = [1, 1, 1], k = 2", []int{1, 1, 1}, 2, 2},
		{"nums = [1, 2, 1, 2, 1], k = 3", []int{1, 2, 1, 2, 1}, 3, 4},
		{"nums = [-1, -1, 1], k = 0", []int{-1, -1, 1}, 0, 1},
		{"nums = [1, -1, 0], k = 0", []int{1, -1, 0}, 0, 3},
		{"nums = [1, 2, 3], k = 3", []int{1, 2, 3}, 3, 2},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectation, subarraySum(tt.nums, tt.k))
		})
	}
}
