package remove_duplicates_from_sorted_array_ii

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want int
	}{
		{"nums = [1,1,1,2,2,3]", []int{1, 1, 1, 2, 2, 3}, 5},
		{"nums = [0,0,1,1,1,1,2,3,3]", []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
		{"nums = [1]", []int{1}, 1},
		{"nums = [0,0,1,1,1,1,2,3,3]", []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, removeDuplicates(tt.nums))
		})
	}
}
