package intersection_of_two_arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	testCases := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{"nums1 = [1,2,2,1], nums2 = [2,2]", []int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{"nums1 = [4,9,5], nums2 = [9,4,9,8,4]", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, intersection(tt.nums1, tt.nums2))
		})
	}
}
