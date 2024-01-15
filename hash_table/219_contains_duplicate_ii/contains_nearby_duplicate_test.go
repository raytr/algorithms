package contains_duplicate_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	testcases := []struct {
		description string
		nums        []int
		k           int
		want        bool
	}{
		{"nums = [1,2,3,1], k = 3", []int{1, 2, 3, 1}, 3, true},
		{"nums = [1,0,1,1], k = 1", []int{1, 0, 1, 1}, 1, true},
		{"nums = [1,2,3,1,2,3], k = 2", []int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for _, tt := range testcases {
		t.Run(tt.description, func(t *testing.T) {
			assert.Equal(t, tt.want, containsNearbyDuplicate(tt.nums, tt.k))
		})
	}
}
