package jump_game

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanJump(t *testing.T) {
	testCases := []struct {
		description string
		nums        []int
		expectation bool
	}{
		{"nums = [2,3,1,1,4]", []int{2, 3, 1, 1, 4}, true},
		// Jump 1 step from index 0 to 1, then 3 steps to the last index.

		{"nums = [3,2,1,0,4]", []int{3, 2, 1, 0, 4}, false},
		// Its maximum jump length is 0, which makes it impossible to reach the last index.
		// You will always arrive at index 3 no matter what.

		{"nums = [1", []int{1}, true},
		{"nums = [0]", []int{0}, true},
		{"nums = [2,0]", []int{2, 0}, true},
	}

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			require.Equal(t, tt.expectation, canJump(tt.nums))
		})
	}
}
