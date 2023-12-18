package climbing_stairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	testcases := []struct {
		description string
		n           int
		want        int
	}{
		{"n = 2", 2, 2},
		{"n = 3", 3, 3},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, climbStairs(tc.n))
		})
	}
}
