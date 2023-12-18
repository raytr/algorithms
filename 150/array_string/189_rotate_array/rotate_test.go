package rotate_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {
	testcases := []struct {
		description string
		nums        []int
		k           int
		want        []int
	}{
		{"nums = [1,2,3,4,5,6,7], k = 3", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{"nums = [-1,-100,3,99], k = 2", []int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{"nums = [1], k = 1", []int{1}, 1, []int{1}},
		{"nums = [1,2], k = 2", []int{1, 2}, 2, []int{1, 2}},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			rotate(tc.nums, tc.k)
			assert.Equal(t, tc.want, tc.nums)
		})
	}
}
