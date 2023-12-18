package move_zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		want []int
	}{
		{"nums = [0,1,0,3,12]", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"nums = [0]", []int{0}, []int{0}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, moveZeroes(tt.nums))
		})
	}
}
