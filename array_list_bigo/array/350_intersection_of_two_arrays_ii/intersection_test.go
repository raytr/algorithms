package intersection_of_two_arrays_ii

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntersection(t *testing.T) {
	testcases := []struct {
		description string
		nums1       []int
		nums2       []int
		expectation []int
	}{
		{"nums1 = [1,2,2,1], nums2 = [2,2]", []int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{"nums1 = [1,2,1,2], nums2 = [2,2]", []int{1, 2, 1, 2}, []int{2, 2}, []int{2, 2}},
		{"nums1 = [4,9,5], nums2 = [9,4,9,8,4]", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			require.ElementsMatch(t, testcase.expectation, intersect(testcase.nums1, testcase.nums2))
		})
	}
}
