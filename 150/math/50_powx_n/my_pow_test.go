package powx_n

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyPow(t *testing.T) {
	testcases := []struct {
		description string
		x           float64
		n           int
		want        float64
	}{
		{"x = 2.00000, n = 10", 2.00000, 10, 1024.00000},
		{"x = 2.10000, n = 3", 2.10000, 3, 9.26100},
		{"x = 2.00000, n = -2", 2.00000, -2, 0.25000},
		{"x = 2.00000, n = 0", 2.00000, 0, 1},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			assert.Equal(t, tc.want, myPow(tc.x, tc.n))
		})
	}
}
