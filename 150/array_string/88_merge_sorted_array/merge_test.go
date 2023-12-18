package merge_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		description string
		nums1       []int
		m           int
		nums2       []int
		n           int
		want        []int
	}{
		{"nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{"nums1 = [1], m = 1, nums2 = [], n = 0", []int{1}, 1, []int{}, 0, []int{1}},
		{"nums1 = [0], m = 0, nums2 = [1], n = 1", []int{0}, 0, []int{1}, 1, []int{1}},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			require.ElementsMatch(t, tt.want, tt.nums1)
		})
	}
}
