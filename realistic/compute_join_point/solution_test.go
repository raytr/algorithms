package compute_join_point

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	tests := []struct {
		name        string
		s1          int
		s2          int
		expectation int
	}{
		{"", 471, 480, 519},
		{"", 11, 7, 620},
		{"", 1158, 2085, 2103},
		{"", 14689, 13, 20014},
		{"", 991, 997, 1118},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectation, ComputeJoinPoint(tt.s1, tt.s2))
		})
	}
}
