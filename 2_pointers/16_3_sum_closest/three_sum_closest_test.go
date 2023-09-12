package sum_closest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestThreeSumClosest(t *testing.T) {
	testCases := []struct {
		name        string
		nums        []int
		target      int
		expectation int
	}{
		{"nums = [-1,2,1,-4], target = 1", []int{-1, 2, 1, -4}, 1, 2},
		{"nums = [0,0,0], target = 1", []int{0, 0, 0}, 1, 0},
		{"nums = [1,1,1,1], target = 3", []int{1, 1, 1, 1}, 3, 3},
		{"nums = [4,0,5,-5,3,3,0,-4,-5], target = -2", []int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2, -2},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectation, threeSumClosest(tt.nums, tt.target))
		})
	}
}
