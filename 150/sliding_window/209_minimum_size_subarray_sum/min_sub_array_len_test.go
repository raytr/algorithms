package minimum_size_subarray_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	testcases := []struct {
		description string
		nums        []int
		target      int
		want        int
	}{
		{"target = 7, nums = [2,3,1,2,4,3]", []int{2, 3, 1, 2, 4, 3}, 7, 2},
		{"target = 4, nums = [1,4,4]", []int{1, 4, 4}, 4, 1},
		{"target = 11, nums = [1,1,1,1,1,1,1,1]", []int{1, 1, 1, 1, 1, 1, 1, 1}, 11, 0},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, minSubArrayLen(tc.target, tc.nums))
		})
	}
}
