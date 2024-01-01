package majority_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	testcases := []struct {
		description string
		nums        []int
		want        int
	}{
		{"nums = [3,2,3]", []int{3, 2, 3}, 3},
		{"nums = [2,2,1,1,1,2,2]", []int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, majorityElement(tc.nums))
		})
	}
}
