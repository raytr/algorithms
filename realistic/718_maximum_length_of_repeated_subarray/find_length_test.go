package maximum_length_of_repeated_subarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLength(t *testing.T) {
	tests := []struct {
		name        string
		nums1       []int
		nums2       []int
		expectation int
	}{
		{"nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]", []int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3},
		{"nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]", []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectation, findLengthBruteForce(tt.nums1, tt.nums2))
		})
	}
}
